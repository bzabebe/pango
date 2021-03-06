package objs

import (
    "testing"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestFwInitialize(t *testing.T) {
    mc := &testdata.MockClient{}
    o := &FwObjs{}
    o.Initialize(mc)

    if o.Address == nil || o.AddressGroup == nil || o.Services == nil || o.ServiceGroup == nil || o.Tags == nil {
        t.Fail()
    }
}

