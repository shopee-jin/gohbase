package gohbase

import (
	"context"
	"io"
	"math/rand"
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

	resultsCh   chan *hrpc.Result
	parallelism int
}

func NewParallelScanner(c *client, rpc *hrpc.Scan, parallel int) hrpc.Scanner {
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
		regions:     regs,
		rootClient:  c,
		rootScan:    rpc,
		resultsCh:   make(chan *hrpc.Result, 10000),
		parallelism: parallel,
	}
}

func (p *parallelScanner) fetch() {
	//we permutate the region first to avoid sychronous behavior
	dest := make([]hrpc.RegionInfo, len(p.regions))
	perm := rand.Perm(len(p.regions))
	for i, v := range perm {
		dest[v] = p.regions[i]
	}

	ch := make(chan hrpc.RegionInfo, 500)
	//put regions into the channel
	go func() {
		for _, region := range dest {
			ch <- region
		}
		close(ch)
	}()

	//now we start workers
	var wg sync.WaitGroup
	for i := 0; i < p.parallelism; i++ {
		wg.Add(1)
		go func() {
			for region := range ch {
				//the worker due with this region
				scanObj, err := hrpc.NewScanRange(p.rootScan.Context(), p.rootScan.Table(), region.StartKey(), region.StopKey(), hrpc.Families(p.rootScan.Families()), hrpc.Filters(p.rootScan.Filter()))
				if err != nil {
					log.Warnf("Create scan object err=%s", err)
					continue
				}
				cli := NewClient(p.rootClient.zkClient.GetQuorum(), p.rootClient.options...)
				sc := cli.Scan(scanObj)
				for {
					//scan one row
					if scanRsp, err := sc.Next(); err != nil {
						//this region finish
						break
					} else {
						p.resultsCh <- scanRsp
					}
				}
				sc.Close()
			}
			wg.Done()
		}()
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
	return nil
}
