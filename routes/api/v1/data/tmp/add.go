package tmp

import (
	"github.com/binganao/Taio/model/response"
	"github.com/binganao/Taio/service/save"
	"github.com/binganao/Taio/utils/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddDataTmp(c *gin.Context) {
	host, _ := c.GetPostForm("host")
	ports, _ := c.GetPostForm("ports")
	services, _ := c.GetPostForm("services")
	fingers, _ := c.GetPostForm("fingers")

	dehost := crypto.Base64Decrypto(host)
	deports := crypto.Base64Decrypto(ports)
	deservices := crypto.Base64Decrypto(services)
	defingers := crypto.Base64Decrypto(fingers)

	if dehost != "" && deports != "" {
		save.TmpSave(dehost, deports, deservices, defingers)

		c.JSON(http.StatusOK, response.Data{
			Code: http.StatusOK,
			Msg:  "数据添加成功",
			Result: response.ProbM{Host: dehost,
				Ports:    deports,
				Services: deservices,
				Fingers:  defingers,
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, response.Data{
			Code: http.StatusBadRequest,
			Msg:  "数据添加失败",
			Result: response.ProbM{Host: dehost,
				Ports:    deports,
				Services: deservices,
				Fingers:  defingers,
			},
		})
	}

}
