package remoteenforcer

import (
	"testing"

	"github.com/headingy/trireme/collector"
	"github.com/headingy/trireme/policy"
)

func TestNewCollector(t *testing.T) {
	Convey("When I create a new collector", t, func() {
		c := NewCollector()
		Convey("The collector should not be nil ", func() {
			So(c, ShouldNotBeNil)
			So(c.Flows, ShouldNotBeNil)
		})
	})
}

func TestCollectFlowEvent(t *testing.T) {
	Convey("Given a stats collector", t, func() {
		c := &CollectorImpl{
			Flows: map[string]*collector.FlowRecord{},
		}

		Convey("When I add a flow event", func() {
			r := &collector.FlowRecord{
				ContextID: "1",
				Source: &collector.EndPoint{
					ID:   "A",
					IP:   "1.1.1.1",
					Type: collector.PU,
				},
				Destination: &collector.EndPoint{
					ID:   "B",
					IP:   "2.2.2.2",
					Type: collector.PU,
					Port: 80,
				},
				Count: 0,
				Tags:  policy.NewTagStore(),
			}
			c.CollectFlowEvent(r)

			Convey("The flow should be in the cache", func() {
				So(len(c.Flows), ShouldEqual, 1)
				So(c.Flows[collector.StatsFlowHash(r)], ShouldNotBeNil)
				So(c.Flows[collector.StatsFlowHash(r)].Count, ShouldEqual, 1)
			})

			Convey("When I add a second flow that matches", func() {
				r := &collector.FlowRecord{
					ContextID: "1",
					Source: &collector.EndPoint{
						ID:   "A",
						IP:   "1.1.1.1",
						Type: collector.PU,
					},
					Destination: &collector.EndPoint{
						ID:   "B",
						IP:   "2.2.2.2",
						Type: collector.PU,
						Port: 80,
					},
					Count: 10,
					Tags:  policy.NewTagStore(),
				}
				c.CollectFlowEvent(r)
				Convey("The flow should be in the cache", func() {
					So(len(c.Flows), ShouldEqual, 1)
					So(c.Flows[collector.StatsFlowHash(r)], ShouldNotBeNil)
					So(c.Flows[collector.StatsFlowHash(r)].Count, ShouldEqual, 11)
				})
			})

			Convey("When I add a third flow that doesn't  matche the previous flows ", func() {
				r := &collector.FlowRecord{
					ContextID: "1",
					Source: &collector.EndPoint{
						ID:   "C",
						IP:   "3.3.3.3",
						Type: collector.PU,
					},
					Destination: &collector.EndPoint{
						ID:   "D",
						IP:   "4.4.4.4",
						Type: collector.PU,
						Port: 80,
					},
					Count: 33,
					Tags:  policy.NewTagStore(),
				}
				c.CollectFlowEvent(r)
				Convey("The flow should be in the cache", func() {
					So(len(c.Flows), ShouldEqual, 2)
					So(c.Flows[collector.StatsFlowHash(r)], ShouldNotBeNil)
					So(c.Flows[collector.StatsFlowHash(r)].Count, ShouldEqual, 33)
				})
			})
		})
	})
}
