package tmp

import (
	"github.com/binganao/TaiO/lib"
	"github.com/binganao/TaiO/model/db"
	"github.com/binganao/TaiO/model/response"
	"github.com/binganao/TaiO/utils/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SingleTmp(c *gin.Context) {
	host := c.Query("host")
	dehost := crypto.Base64Decrypto(host)
	if dehost != "" {
		c.JSON(http.StatusOK, queryHost(dehost))
	} else {
		c.JSON(http.StatusBadRequest, response.Search{Code: http.StatusBadRequest})
	}
}

func queryHost(host string) response.Search {
	s := response.Search{}
	var qd []db.ProbTmpM

	lib.GetDB().Where("host = ?", host).Find(&qd)

	if len(qd) > 0 {
		s.Code = http.StatusOK
		s.Msg = "操作成功"
		s.Hosts = qd[0].Host
		ports := strings.Split(qd[0].Ports, ",")
		s.Ports = ports
		var ser []response.Service
		se := strings.Split(qd[0].Services, ",")
		for _, ss := range se {
			if strings.Contains(ss, ":") {
				part := strings.Split(ss, ":")
				sv := response.Service{}
				sv.Port = part[0]
				sv.ServiceName = part[1]
				snk := map[string]string{
					"ftp":           "FTP",
					"mysql":         "MYSQL",
					"ssh":           "SSH",
					"vnc":           "VNC",
					"ms-wbt-server": "RDP",
					"ms-sql-s":      "MSSQL",
					"redis":         "REDIS",
				}
				sv.MaybeDanger = false
				for key, value := range snk {
					if strings.Contains(sv.ServiceName, key) {
						sv.ServiceNick = value
						sv.MaybeDanger = true
					}
				}
				ser = append(ser, sv)
			}
		}
		s.Services = ser

		var ftss []response.Finger
		fts := strings.Split(qd[0].Fingers, ";")
		for _, ft := range fts {
			if strings.Contains(ft, ",") {
				part := strings.Split(ft, ",")
				f := response.Finger{}
				f.Url = part[0]
				f.App = part[1]
				f.Server = part[2]
				ftss = append(ftss, f)
			}
		}
		s.Fingers = ftss

		return s
	} else {
		return response.Search{}
	}
}
