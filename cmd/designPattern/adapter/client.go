package main

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	com.InsertIntoLightningPort()
}
