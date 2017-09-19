// Copyright (C) 2017  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package gohbase

import (
	"context"
	"io"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/jasonzzw/gohbase/filter"
	"github.com/jasonzzw/gohbase/hrpc"
	"github.com/jasonzzw/gohbase/region"
)

type parallelScanner struct {
	regions    []hrpc.RegionInfo
	rootClient *client
	rootScan   *hrpc.Scan
	once       sync.Once

	subScanner []hrpc.Scanner
	resultsCh  chan *hrpc.Result
}

func NewParallelScanner(c *client, rpc *hrpc.Scan) hrpc.Scanner {
	regs := []hrpc.RegionInfo{}

	var scanRequest *hrpc.Scan
	var err error
	scanRequest, err = hrpc.NewScanStr(context.Background(), "hbase:meta",
		hrpc.Filters(filter.NewPrefixFilter([]byte(rpc.Table()))))
	if err != nil {
		log.Errorf("Error with scanning hbase:meta table, err=%s", err)
		return nil
	}
	//else, we scan the table info
	scanner := c.Scan(scanRequest)
	for {
		//scan one row
		if scanRsp, err := scanner.Next(); err != nil {
			break
		} else {
			if regionInfo, _, err := region.ParseRegionInfo(scanRsp); err == nil {
				regs = append(regs, regionInfo)
			}
		}
	}
	log.Infof("Retrieve %d regions for table=%s", len(regs), rpc.Table())

	return &parallelScanner{
		regions:    regs,
		rootClient: c,
		rootScan:   rpc,
		resultsCh:  make(chan *hrpc.Result, 10000),
	}
}

func (p *parallelScanner) fetch() {
	var wg sync.WaitGroup
	for _, region := range p.regions {
		scanObj, err := hrpc.NewScanRange(p.rootScan.Context(), p.rootScan.Table(), region.StartKey(), region.StopKey(), hrpc.Families(p.rootScan.Families()), hrpc.Filters(p.rootScan.Filter()))
		if err != nil {
			log.Warnf("Create scan object err=%s", err)
			continue
		}
		cli := NewClient(p.rootClient.zkClient.GetQuorum(), p.rootClient.options...)
		sc := cli.Scan(scanObj)
		//save the scanner object for canceling
		p.subScanner = append(p.subScanner, sc)
		wg.Add(1)
		go func(sc hrpc.Scanner) {
			for {
				//scan one row
				if scanRsp, err := sc.Next(); err != nil {
					break
				} else {
					p.resultsCh <- scanRsp
				}
			}
			sc.Close()
			wg.Done()
		}(sc)
	}
	wg.Wait()
	close(p.resultsCh)
}

func (p *parallelScanner) Next() (*hrpc.Result, error) {
	p.once.Do(func() {
		go p.fetch()
	})

	//return nil, nil
	if result, ok := <-p.resultsCh; ok {
		return result, nil
	} else {
		return nil, io.EOF
	}
}

func (p *parallelScanner) Close() error {
	for _, sc := range p.subScanner {
		sc.Close()
	}
	return nil
}
