<h2>Open Source Wethenew Monitor Based on GoLang</h2>

_A Go Web Application to scrape and monitoring an API_

<h3>The module contains the following Features</h3>

- Ability to send a [HTTPS](https://pkg.go.dev/net/http) request to Wethenew API every seconds
- Supporting [TLS](https://tls13.xargs.org/) client

```go

http "github.com/bogdanfinn/fhttp"
tls_client "github.com/bogdanfinn/tls-client"

options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(8),
		tls_client.WithClientProfile(tls_client.Chrome_105),
		tls_client.WithInsecureSkipVerify(),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
    if err != nil {
        log.Fatal(err)
    }
```

- Comparing the old items with the new-one to send

```go
for idx, v := range new_id.Results {
    if !Contains(Slice, v.ID) {
        Slice = append(Slice, v.ID)
        discord.Webhook(new_id, idx)
        continue
    }
}
```

- Ability to send the new items to the client via Webhook message to Discord

```go
payload := &data.Top{
		Username:  "Wethenew Monitor",
		AvatarURL: Image_URL,
		Content:   "",
		Embeds: []data.Embeds{
			{
				Title: new_id.Results[idx].Name,
				Color:  1999236,
				Fields: fields,
				Thumbnail: data.Thumbnail{
					URL: new_id.Results[idx].Image,
				},
				Footer: data.Footer{
					IconURL: Image_URL,
					Text:    "Wethenew | Holding-Lab " + time,
				},
			},
		},
	}
```

- Auto-restart the application when the application is crashed

## Deployment

- Run the web application on Docker

[Go Web Application with Docker](https://index.docker.io/_/golang)

```bash
docker build -t wethenew-monitor .
docker run -it --rm --name <app_name> wethenew-monitor
```

## Contributing

To contribute follow these steps:

1. Fork this repository.
2. Create a branch with clear name: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin <project_name>/<location>`
5. Create the pull request.

Alternatively see the GitHub documentation on [creating a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).
