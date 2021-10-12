package client

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"gopkg.in/guregu/null.v4"
	"io"
	"strings"
)

type DashboardWidget struct {
	GUID                        null.String `json:"guid"`
	Label                       null.String `json:"label"`
	Title                       null.String `json:"title"`
	Type                        null.String `json:"type"`
	Height                      null.Int    `json:"height"`
	Width                       null.Int    `json:"width"`
	X                           null.Int    `json:"x"`
	Y                           null.Int    `json:"y"`
	AdqlQueries                 []string    `json:"adqlQueries"`
	AnalyticsType               null.String `json:"analyticsType"`
	SearchMode                  null.String `json:"searchMode"`
	WidgetsMetricMatchCriterias null.String `json:"widgetsMetricMatchCriterias"`
	Description                 null.String `json:"description"`
	DrillDownUrl                null.String `json:"drillDownUrl"`
	UseMetricBrowserAsDrillDown null.Bool   `json:"useMetricBrowserAsDrillDown"`
	DrillDownActionType         null.String `json:"drillDownActionType"`
	BackgroundColor             null.Int    `json:"backgroundColor"`
	Color                       null.Int    `json:"color"`
	FontSize                    null.Int    `json:"fontSize"`
	UseAutomaticFontSize        null.Bool   `json:"useAutomaticFontSize"`
	BorderEnabled               null.Bool   `json:"borderEnabled"`
	BorderThickness             null.Int    `json:"borderThickness"`
	BorderColor                 null.Int    `json:"borderColor"`
	BackgroundAlpha             null.Float  `json:"backgroundAlpha"`
	ShowValues                  null.Bool   `json:"showValues"`
	FormatNumber                null.Bool   `json:"formatNumber"`
	NumDecimals                 null.Int    `json:"numDecimals"`
	RemoveZeros                 null.Bool   `json:"removeZeros"`
	BackgroundColors            []int       `json:"backgroundColors"`
	CompactMode                 null.Bool   `json:"compactMode"`
	ShowTimeRange               null.Bool   `json:"showTimeRange"`
	RenderIn3D                  null.Bool   `json:"renderIn3D"`
	ShowLegend                  null.Bool   `json:"showLegend"`
	LegendPosition              null.String `json:"legendPosition"`
	LegendColumnCount           null.Int    `json:"legendColumnCount"`
	StartTime                   null.Bool   `json:"startTime"`       // null
	EndTime                     null.Int    `json:"endTime"`         // null
	CustomTimeRange             null.Int    `json:"customTimeRange"` // null ?
	MinutesBeforeAnchorTime     null.Int    `json:"minutesBeforeAnchorTime"`
	IsGlobal                    null.Bool   `json:"isGlobal"`

	// TODO
	//Properties          []int  `json:"properties"` // null list?
	//MissingEntities     null.String `json:"missingEntities"`
	//VerticalAxisLabel   null.String `json:"verticalAxisLabel"`
	//HideHorizontalAxis  null.String `json:"hideHorizontalAxis"`
	//HorizontalAxisLabel null.String `json:"horizontalAxisLabel"`
	//AxisType string `json:"axisType"`
	//MultipleYAxis         bool   `json:"multipleYAxis"`
	//CustomVerticalAxisMin bool   `json:"customVerticalAxisMin"`
	//CustomVerticalAxisMax bool   `json:"customVerticalAxisMax"`
	//ShowEvents            null.Bool   `json:"showEvents"`
	//EventFilter           bool   `json:"eventFilter"`
	//InterpolateDataGaps bool `json:"interpolateDataGaps"`
	//ShowAllTooltips       null.Bool   `json:"showAllTooltips"`
	//StaticThresholds      null.String `json:"staticThresholds"`

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
	return fmt.Sprintf("%s/dashboardIfUpdated/%d/-1", c.createDashboardBaseUrl(), dashboardId)
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

	updated := Dashboard{
		Name: dashboard.Name,
		ID:   dResponse.Dashboard.ID,
	}
	return &updated, err
}

func (c *AppDClient) CreateDashboard(dashboard Dashboard) (*Dashboard, error) {
	resp, err := req.Post(c.createDashboardUrl(), c.createAuthHeader(), req.BodyJSON(dashboard))
	if resp.Response().StatusCode != 200 {
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

func (c *AppDClient) UpdateDashboard(dashboard Dashboard) (*Dashboard, error) {
	resp, err := req.Post(c.updateDashboardUrl(), c.createAuthHeader(), req.BodyJSON(dashboard))
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error updating Dashboard: %d, %s", resp.Response().StatusCode, respString))

	}
	updated := Dashboard{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, err
}

func (c *AppDClient) DeleteDashboard(dashboardId int) error {
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
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error getting dashboard: %d, %s", resp.Response().StatusCode, respString))
	}
	dashboard := Dashboard{}
	err = resp.ToJSON(&dashboard)

	if err != nil {
		return nil, err
	}
	dashboard.ID = dashboardId
	return &dashboard, nil
}
