package bssopenapi

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

// Item is a nested struct in bssopenapi response
type Item struct {
	Tax                       float64                `json:"Tax" xml:"Tax"`
	CoveragePercentage        float64                `json:"CoveragePercentage" xml:"CoveragePercentage"`
	ResourceInstanceId        string                 `json:"ResourceInstanceId" xml:"ResourceInstanceId"`
	CommodityCode             string                 `json:"CommodityCode" xml:"CommodityCode"`
	BudgetStatus              string                 `json:"BudgetStatus" xml:"BudgetStatus"`
	ConsumePeriod             map[string]interface{} `json:"ConsumePeriod" xml:"ConsumePeriod"`
	PipCode                   string                 `json:"PipCode" xml:"PipCode"`
	SplitAccountID            string                 `json:"SplitAccountID" xml:"SplitAccountID"`
	BillAccountID             string                 `json:"BillAccountID" xml:"BillAccountID"`
	ListPrice                 string                 `json:"ListPrice" xml:"ListPrice"`
	Period                    string                 `json:"Period" xml:"Period"`
	InvoiceDiscount           float64                `json:"InvoiceDiscount" xml:"InvoiceDiscount"`
	StartTime                 string                 `json:"StartTime" xml:"StartTime"`
	EndTime                   string                 `json:"EndTime" xml:"EndTime"`
	InstanceID                string                 `json:"InstanceID" xml:"InstanceID"`
	PretaxGrossAmount         float64                `json:"PretaxGrossAmount" xml:"PretaxGrossAmount"`
	ImageType                 string                 `json:"ImageType" xml:"ImageType"`
	RecordID                  string                 `json:"RecordID" xml:"RecordID"`
	Status                    string                 `json:"Status" xml:"Status"`
	Item                      string                 `json:"Item" xml:"Item"`
	ProductName               string                 `json:"ProductName" xml:"ProductName"`
	PaymentAmount             float64                `json:"PaymentAmount" xml:"PaymentAmount"`
	PostpaidCost              string                 `json:"PostpaidCost" xml:"PostpaidCost"`
	ReservationCost           string                 `json:"ReservationCost" xml:"ReservationCost"`
	UsageEndTime              string                 `json:"UsageEndTime" xml:"UsageEndTime"`
	InternetIP                string                 `json:"InternetIP" xml:"InternetIP"`
	CashAmount                float64                `json:"CashAmount" xml:"CashAmount"`
	ResourceGroup             string                 `json:"ResourceGroup" xml:"ResourceGroup"`
	SplitAccountName          string                 `json:"SplitAccountName" xml:"SplitAccountName"`
	CapacityUnit              string                 `json:"CapacityUnit" xml:"CapacityUnit"`
	UsagePercentage           float64                `json:"UsagePercentage" xml:"UsagePercentage"`
	SubscribeTime             string                 `json:"SubscribeTime" xml:"SubscribeTime"`
	InstanceId                string                 `json:"InstanceId" xml:"InstanceId"`
	SubscriptionType          string                 `json:"SubscriptionType" xml:"SubscriptionType"`
	ZoneName                  string                 `json:"ZoneName" xml:"ZoneName"`
	PotentialSavedCost        string                 `json:"PotentialSavedCost" xml:"PotentialSavedCost"`
	MultAccountRelSubscribe   string                 `json:"MultAccountRelSubscribe" xml:"MultAccountRelSubscribe"`
	Percentage                float64                `json:"Percentage" xml:"Percentage"`
	SplitCommodityCode        string                 `json:"SplitCommodityCode" xml:"SplitCommodityCode"`
	SplitItemID               string                 `json:"SplitItemID" xml:"SplitItemID"`
	SubscribeType             string                 `json:"SubscribeType" xml:"SubscribeType"`
	AdjustAmount              float64                `json:"AdjustAmount" xml:"AdjustAmount"`
	ProductType               string                 `json:"ProductType" xml:"ProductType"`
	SplitBillingDate          string                 `json:"SplitBillingDate" xml:"SplitBillingDate"`
	PaymentTransactionID      string                 `json:"PaymentTransactionID" xml:"PaymentTransactionID"`
	BillingDate               string                 `json:"BillingDate" xml:"BillingDate"`
	StatusName                string                 `json:"StatusName" xml:"StatusName"`
	BillingItem               string                 `json:"BillingItem" xml:"BillingItem"`
	SavedCost                 string                 `json:"SavedCost" xml:"SavedCost"`
	DeductQuantity            float64                `json:"DeductQuantity" xml:"DeductQuantity"`
	RowLimitPerFile           int                    `json:"RowLimitPerFile" xml:"RowLimitPerFile"`
	BudgetName                string                 `json:"BudgetName" xml:"BudgetName"`
	Budget                    map[string]interface{} `json:"Budget" xml:"Budget"`
	ItemName                  string                 `json:"ItemName" xml:"ItemName"`
	TotalQuantity             float64                `json:"TotalQuantity" xml:"TotalQuantity"`
	BizType                   string                 `json:"BizType" xml:"BizType"`
	BillingItemCode           string                 `json:"BillingItemCode" xml:"BillingItemCode"`
	BucketPath                string                 `json:"BucketPath" xml:"BucketPath"`
	UsageStartTime            string                 `json:"UsageStartTime" xml:"UsageStartTime"`
	ProductDetail             string                 `json:"ProductDetail" xml:"ProductDetail"`
	Usage                     string                 `json:"Usage" xml:"Usage"`
	ServicePeriodUnit         string                 `json:"ServicePeriodUnit" xml:"ServicePeriodUnit"`
	PretaxAmountLocal         float64                `json:"PretaxAmountLocal" xml:"PretaxAmountLocal"`
	OwnerName                 string                 `json:"OwnerName" xml:"OwnerName"`
	OutstandingAmount         float64                `json:"OutstandingAmount" xml:"OutstandingAmount"`
	DeductedByResourcePackage string                 `json:"DeductedByResourcePackage" xml:"DeductedByResourcePackage"`
	SplitProductDetail        string                 `json:"SplitProductDetail" xml:"SplitProductDetail"`
	ProductCode               string                 `json:"ProductCode" xml:"ProductCode"`
	BucketOwnerId             int64                  `json:"BucketOwnerId" xml:"BucketOwnerId"`
	SplitItemName             string                 `json:"SplitItemName" xml:"SplitItemName"`
	SubscribeBucket           string                 `json:"SubscribeBucket" xml:"SubscribeBucket"`
	Region                    string                 `json:"Region" xml:"Region"`
	RoundDownDiscount         string                 `json:"RoundDownDiscount" xml:"RoundDownDiscount"`
	CommodityName             string                 `json:"CommodityName" xml:"CommodityName"`
	PaymentTime               string                 `json:"PaymentTime" xml:"PaymentTime"`
	CalculatedValues          map[string]interface{} `json:"CalculatedValues" xml:"CalculatedValues"`
	CostUnit                  string                 `json:"CostUnit" xml:"CostUnit"`
	SplitBillingCycle         string                 `json:"SplitBillingCycle" xml:"SplitBillingCycle"`
	AfterTaxAmount            float64                `json:"AfterTaxAmount" xml:"AfterTaxAmount"`
	SubOrderId                string                 `json:"SubOrderId" xml:"SubOrderId"`
	InstanceSpec              string                 `json:"InstanceSpec" xml:"InstanceSpec"`
	BillingType               string                 `json:"BillingType" xml:"BillingType"`
	Tag                       string                 `json:"Tag" xml:"Tag"`
	RegionNo                  string                 `json:"RegionNo" xml:"RegionNo"`
	OwnerID                   string                 `json:"OwnerID" xml:"OwnerID"`
	UserId                    string                 `json:"UserId" xml:"UserId"`
	NickName                  string                 `json:"NickName" xml:"NickName"`
	InstanceConfig            string                 `json:"InstanceConfig" xml:"InstanceConfig"`
	DeductedByCashCoupons     float64                `json:"DeductedByCashCoupons" xml:"DeductedByCashCoupons"`
	ServicePeriod             string                 `json:"ServicePeriod" xml:"ServicePeriod"`
	ListPriceUnit             string                 `json:"ListPriceUnit" xml:"ListPriceUnit"`
	UsageUnit                 string                 `json:"UsageUnit" xml:"UsageUnit"`
	PaymentCurrency           string                 `json:"PaymentCurrency" xml:"PaymentCurrency"`
	SubscribeLanguage         string                 `json:"SubscribeLanguage" xml:"SubscribeLanguage"`
	BudgetType                string                 `json:"BudgetType" xml:"BudgetType"`
	Currency                  string                 `json:"Currency" xml:"Currency"`
	DeductedByPrepaidCard     float64                `json:"DeductedByPrepaidCard" xml:"DeductedByPrepaidCard"`
	BillAccountName           string                 `json:"BillAccountName" xml:"BillAccountName"`
	DeductedByCoupons         float64                `json:"DeductedByCoupons" xml:"DeductedByCoupons"`
	Zone                      string                 `json:"Zone" xml:"Zone"`
	Quantity                  int64                  `json:"Quantity" xml:"Quantity"`
	UserName                  string                 `json:"UserName" xml:"UserName"`
	PretaxAmount              float64                `json:"PretaxAmount" xml:"PretaxAmount"`
	IntranetIP                string                 `json:"IntranetIP" xml:"IntranetIP"`
}
