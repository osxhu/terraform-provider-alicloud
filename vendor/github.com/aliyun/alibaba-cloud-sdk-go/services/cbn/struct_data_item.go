package cbn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DataItem is a nested struct in cbn response
type DataItem struct {
	Name                             string `json:"Name" xml:"Name"`
	Nexthop                          string `json:"Nexthop" xml:"Nexthop"`
	TrRouteTableId                   string `json:"TrRouteTableId" xml:"TrRouteTableId"`
	InstanceId                       string `json:"InstanceId" xml:"InstanceId"`
	RouteType                        string `json:"RouteType" xml:"RouteType"`
	TransitRouteTableAggregationCidr string `json:"TransitRouteTableAggregationCidr" xml:"TransitRouteTableAggregationCidr"`
	Scope                            string `json:"Scope" xml:"Scope"`
	Status                           string `json:"Status" xml:"Status"`
	Description                      string `json:"Description" xml:"Description"`
}
