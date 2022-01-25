package tmp

import (
	"github.com/binganao/TaiO/model/response"
	"github.com/binganao/TaiO/service/datas"
	"github.com/binganao/TaiO/utils/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddDataTmp(c *gin.Context) {
	host, _ := c.GetPostForm("host")
	ports, _ := c.GetPostForm("ports")
	services, _ := c.GetPostForm("services")
	fingers, _ := c.GetPostForm("fingers")

	deHost := crypto.Base64Decrypto(host)
	dePorts := crypto.Base64Decrypto(ports)
	deServices := crypto.Base64Decrypto(services)
	deFingers := crypto.Base64Decrypto(fingers)

	if deHost != "" && dePorts != "" {
		datas.TmpSave(deHost, dePorts, deServices, deFingers)

		c.JSON(http.StatusOK, response.Data{
			Code: http.StatusOK,
			Msg:  "数据添加成功",
			Result: response.ProbM{Host: deHost,
				Ports:    dePorts,
				Services: deServices,
				Fingers:  deFingers,
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, response.Data{
			Code: http.StatusBadRequest,
			Msg:  "数据添加失败",
			Result: response.ProbM{Host: deHost,
				Ports:    dePorts,
				Services: deServices,
				Fingers:  deFingers,
			},
		})
	}

}
