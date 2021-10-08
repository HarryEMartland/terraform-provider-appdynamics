package client

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"io"
	"strings"
)

//"description" : null,
//"drillDownUrl" : null,
//"useMetricBrowserAsDrillDown" : true,
//"drillDownActionType" : null,
//"backgroundColor" : 16777215,
//"color" : 1646891,
//"fontSize" : 12,
//"useAutomaticFontSize" : false,
//"borderEnabled" : false,
//"borderThickness" : 0,
//"borderColor" : 14408667,
//"backgroundAlpha" : 1.0,
//"showValues" : false,
//"formatNumber" : null,
//"numDecimals" : 0,
//"removeZeros" : null,
//"backgroundColors" : [ 16777215, 16777215 ],
//"compactMode" : false,
//"showTimeRange" : false,
//"renderIn3D" : false,
//"showLegend" : true,
//"legendPosition" : "POSITION_BOTTOM",
//"legendColumnCount" : 1,
//"startTime" : null,
//"endTime" : null,
//"customTimeRange" : null,
//"minutesBeforeAnchorTime" : 15,
//"isGlobal" : true,
//"properties" : [ ],
//"missingEntities" : null,
//"verticalAxisLabel" : null,
//"hideHorizontalAxis" : null,
//"horizontalAxisLabel" : null,
//"axisType" : "LINEAR",
//"stackMode" : null,
//"multipleYAxis" : null,
//"customVerticalAxisMin" : null,
//"customVerticalAxisMax" : null,
//"showEvents" : null,
//"eventFilter" : null,
//"interpolateDataGaps" : false,
//"showAllTooltips" : null,
//"staticThresholds" : null

//"id" : 10401366,
//"version" : 0,
//"guid" : "2fce8ed0-2e3b-4cd4-8425-3ff67855d2c0",
//"title" : null,
//"type" : "TIMESERIES_GRAPH",
//"dashboardId" : 3768,
//"widgetsMetricMatchCriterias" : null,
//"height" : 3,
//"width" : 6,
//"minHeight" : 0,
//"minWidth" : 0,
//"x" : 0,
//"y" : 0,
//"label" : null,

type DashboardWidget struct {

	//DashboardId int `json:"dashboardId"`
	//ID                          int    `json:"id"`      // AUTO
	//Version                     int    `json:"version"` // AUTO
	GUID string `json:"guid"`
	//Title                       string `json:"title"`
	Type        string `json:"type"`
	DashboardId int    `json:"dashboardId"` // AUTO
	//WidgetsMetricMatchCriterias string `json:"widgetsMetricMatchCriterias"`
	Height    int `json:"height"`
	Width     int `json:"width"`
	MinHeight int `json:"minHeight"`
	MinWidth  int `json:"minWidth"`
	X         int `json:"x"`
	Y         int `json:"y"`
	//Label                       string `json:"label"`
	//Description                 string `json:"description"`
	//DrillDownUrl                string `json:"drillDownUrl"`
	//BackgroundColor             int    `json:"backgroundColor"`
	//Color                       int    `json:"color"`
	//FontSize                    int    `json:"fontSize"`
	//UseAutomaticFontSize        int    `json:"useAutomaticFontSize"`
	//BorderEnabled               bool   `json:"borderEnabled"`
	//BorderThickness             int    `json:"borderThickness"`
	//BorderColor                 int    `json:"borderColor"`
	//BackgroundAlpha             int    `json:"backgroundAlpha"`
	//ShowValues                  bool   `json:"showValues"`
	//FormatNumber                int    `json:"formatNumber"` // null
	//NumDecimals                 int    `json:"numDecimals"`
	//RemoveZeros                 int    `json:"removeZeros"`
	//BackgroundColors            []int  `json:"backgroundColors"`
	//CompactMode                 bool   `json:"compactMode"`
	//ShowTimeRange               bool   `json:"showTimeRange"`
	//RenderIn3D                  bool   `json:"renderIn3D"`
	//ShowLegend                  bool   `json:"showLegend"`
	//LegendPosition              string `json:"legendPosition"`
	//LegendColumnCount           int    `json:"legendColumnCount"`
	//StartTime                   bool   `json:"startTime"`       // null
	//EndTime                     int    `json:"endTime"`         // null
	//CustomTimeRange             int    `json:"customTimeRange"` // null ?
	//MinutesBeforeAnchorTime     int    `json:"minutesBeforeAnchorTime"`
	//IsGlobal                    int    `json:"isGlobal"`
	//properties     int    `json:"properties"` // null list?
	//MissingEntities       string `json:"missingEntities"`
	//VerticalAxisLabel     string `json:"verticalAxisLabel"`
	//HideHorizontalAxis    string `json:"hideHorizontalAxis"`
	//HorizontalAxisLabel   string `json:"horizontalAxisLabel"`
	//AxisType              string `json:"axisType"`
	//MultipleYAxis         bool   `json:"multipleYAxis"`
	//CustomVerticalAxisMin bool   `json:"customVerticalAxisMin"`
	//CustomVerticalAxisMax bool   `json:"customVerticalAxisMax"`
	//ShowEvents            bool   `json:"showEvents"`
	//EventFilter           bool   `json:"eventFilter"`
	//InterpolateDataGaps   bool   `json:"interpolateDataGaps"`
	//ShowAllTooltips       bool   `json:"showAllTooltips"`
	//StaticThresholds      string `json:"staticThresholds"`
}

type Dashboard struct {
	ID                      int               `json:"id"`
	Name                    string            `json:"name"`
	Width                   int               `json:"width"`
	Height                  int               `json:"height"`
	CanvasType              string            `json:"canvasType"`
	TemplateEntityType      string            `json:"templateEntityType"`
	RefreshInterval         int               `json:"refreshInterval"`
	BackgroundColor         int               `json:"backgroundColor"`
	WarRoom                 bool              `json:"warRoom"`
	Template                bool              `json:"template"`
	Widgets                 []DashboardWidget `json:"widgets"`
	Version                 int               `json:"version"`
	MinutesBeforeAnchorTime int               `json:"minutesBeforeAnchorTime"`
	StartTime               int               `json:"startTime"`
	EndTime                 int               `json:"endTime"`
}

type AppdDashboardCreateResponse struct {
	Success   bool `json:"success"`
	Dashboard struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"dashboard"`
}

func (c *AppDClient) createDashboardBaseUrl() string {
	return fmt.Sprintf("%s/controller/restui/dashboards", c.BaseUrl)
}

//https://worldremit-test.saas.appdynamics.com/controller/restui/dashboards/deleteDashboard
//
func (c *AppDClient) createDashboardUrl() string {
	return fmt.Sprintf("%s/createDashboard", c.createDashboardBaseUrl())
}

func (c *AppDClient) updateDashboardUrl() string {
	return fmt.Sprintf("%s/updateDashboard", c.createDashboardBaseUrl())
}

func (c *AppDClient) deleteDashboard() string {
	return fmt.Sprintf("%s/deleteDashboard", c.createDashboardBaseUrl())
}

func (c *AppDClient) deleteDashboards() string {
	return fmt.Sprintf("%s/deleteDashboards", c.createDashboardBaseUrl())
}

func (c *AppDClient) getDashboardUrl(dashboardId int) string {
	return fmt.Sprintf("%s/controller/CustomDashboardImportExportServlet?dashboardId=%d", c.BaseUrl, dashboardId)
}

func (c *AppDClient) importDashboardUrl() string {
	return fmt.Sprintf("%s/controller/CustomDashboardImportExportServlet", c.BaseUrl)
}

func (c *AppDClient) ImportDashboard(dashboard Dashboard, template string) (*Dashboard, error) {
	file := io.NopCloser(strings.NewReader(template))
	resp, err := req.Post(c.createDashboardUrl(), c.createAuthHeader(), req.FileUpload{
		File:      file,
		FieldName: "fileUpload",    // FieldName is form field name
		FileName:  "template.json", //Filename is the name of the file that you wish to upload. We use this to guess the mimetype as well as pass it onto the server
	})
	dResponse := AppdDashboardCreateResponse{}
	err = resp.ToJSON(&dResponse)
	fmt.Println(dResponse.Dashboard.ID)

	updated := Dashboard{
		Name: dashboard.Name,
		ID:   dResponse.Dashboard.ID,
	}
	return &updated, err
}

func (c *AppDClient) CreateDashboard(dashboard Dashboard) (*Dashboard, error) {

	resp, err := req.Post(c.createDashboardUrl(), c.createAuthHeader(), req.BodyJSON(dashboard))
	fmt.Println("TRYING TO CREATE", dashboard)
	if resp.Response().StatusCode != 201 {
		respString, _ := resp.ToString()
		fmt.Println(resp)
		fmt.Println("XDD")
		return nil, errors.New(fmt.Sprintf("Error creating Dashboard: %d, %s", resp.Response().StatusCode, respString))
	}
	fmt.Println("A")
	updated := Dashboard{}
	fmt.Println("B", resp)
	fmt.Println("BA", resp.Response().StatusCode)
	err = resp.ToJSON(&updated)
	fmt.Println("C", err)
	if err != nil {
		return nil, err
	}
	fmt.Println("D")
	fmt.Println("Result", resp)
	return &updated, err
}

func (c *AppDClient) UpdateDashboard(dashboard Dashboard) (*Dashboard, error) {
	resp, err := req.Post(c.updateDashboardUrl(), c.createAuthHeader(), req.BodyJSON(dashboard))
	if resp.Response().StatusCode != 201 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error creating Dashboard: %d, %s", resp.Response().StatusCode, respString))

	}
	updated := Dashboard{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}
	return &updated, err
}

func (c *AppDClient) DeleteDashboard(dashboardId int) error {
	fmt.Println("REMOVING", dashboardId)
	resp, err := req.Post(c.deleteDashboard(), c.createAuthHeader(), req.BodyJSON(dashboardId))
	if err != nil {
		return err
	}

	if resp.Response().StatusCode != 204 {
		respString, _ := resp.ToString()
		return errors.New(fmt.Sprintf("Error deleting Dashboard: %d, %s", resp.Response().StatusCode, respString))
	}

	return nil
}

func (c *AppDClient) GetDashboard(dashboardId int) (*Dashboard, error) {
	resp, err := req.Get(c.getDashboardUrl(dashboardId), c.createAuthHeader())
	if err != nil {
		return nil, err
	}
	dashboard := Dashboard{}
	err = resp.ToJSON(&dashboard)
	if err != nil {
		return nil, err
	}
	dashboard.ID = dashboardId
	return &dashboard, nil
}
