package thumbnail

import (
  "fmt"
  "github.com/matrix-org/go-neb/types"
  "maunium.net/go/mautrix/id"
  "regexp"
  "strings"
)

const ServiceType = "thumbnail"

var linkRegex = regexp.MustCompile("http[^\\s]+")

type Service struct {
  types.DefaultService
}


func getExtension(s string) string {
    ss := strings.Split(s, ".")
    // There may be query parameters in the end e.g. .gif?format=gif
    lastPart := ss[len(ss)-1]
    return strings.Split(lastPart, "?")[0]
}

func (s *Service) Expansions(client types.MatrixClient) []types.Expansion {
  return []types.Expansion{
    {
      Regexp: linkRegex,
      Expand: func(roomID id.RoomID, userID id.UserID, links []string) interface{} {
        /*
        // Problem: s.serviceUserID is not set.
        if s.serviceUserID == userID {
          fmt.Println("Not replying to myself")
          return nil
        }
        */
        extension := strings.ToLower(getExtension(links[0]))

        switch extension {
        case "png", "jpg", "jpeg", "gif":
          break
        default:
          return nil
        }
        /*
        resUpload, err := client.UploadLink(links[0])
        if err != nil {
          return nil
        }

        return mevt.MessageEventContent{
          MsgType: "m.image",
          Body:    links[0],
          URL:     resUpload.ContentURI.CUString(),
          Info: &mevt.FileInfo{
            MimeType: "image/jpeg",
          },
        }
        */
        return nil
      },
    },
  }
}

func init() {
  types.RegisterService(func(serviceID string, serviceUserID id.UserID, webhookEndpointURL string) types.Service {
    // Problem: serviceUserID is nil or an empty string
    fmt.Println("Initializing service, serviceUserID: " + serviceUserID)
    return &Service{
      DefaultService: types.NewDefaultService(serviceID, serviceUserID, ServiceType),
    }
  })
}
