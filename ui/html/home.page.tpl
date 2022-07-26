{{template "base" .}}

{{define "title"}}Socket io demo{{end}}

{{define "main"}}
    <div class="row">
        <div class="col">
            <div class="card card-body">
                <h5 class="card-title">Message log
                    <button id="clear-log" class="btn btn-warning">clear</button>
                </h5>
                <div id="chat" class=""></div>
            </div>
        </div>
        <div class="col">
            <div class="card card-body">
                <div>
                    <h5>Create new user</h5>
                    <form id="add-user">
                        <table class="table">
                            <tr>
                                <td>Json attributes:</td>
                                <td><textarea class="form-control form-control-sm" cols="10" rows="5"
                                              id="attributes">
{
  "gender": "male",
  "country": "USA",
  "games": ["slots","poker"]
}
                                </textarea></td>
                            </tr>
                            <tr>
                                <td>Callback url</td>
                                <td><input type="text" id="callbackUrl" value=""/></td>
                            </tr>
                            <tr>
                                <td></td>
                                <td><input class="btn btn-success" style="float:right;" type="submit"></td>
                            </tr>
                        </table>
                    </form>
                </div>
                <div>
                    <h5>Room</h5>
                    <form>
                        <table class="table">
                            <tr>
                                <td>Realm</td>
                                <td><input type="text" value="realmName" id="realm"></td>
                            </tr>
                            <tr>
                                <td>room</td>
                                <td><input type="text" value="room1" id="room"></td>
                            </tr>
                            <tr>
                                <td><input class="btn btn-success" style="float:right;" value="Join" id="join"></td>
                                <td><input class="btn btn-success" style="float:right;" value="Leave" id="leave"></td>
                                <td><input class="btn btn-success" style="float:right;" value="RoomInfo" id="info"></td>
                            </tr>
                        </table>
                    </form>
                </div>
                <h4>Send message via API</h4>
                <form id="send-message">
                    <table class="table">
                        <tr>
                            <td>Query:</td>
                            <td><textarea class="form-control form-control-sm" cols="10" rows="5" id="query">
{
  "gender": "male",
  "country": "USA"
}
                            </textarea>
                            </td>
                        </tr>
                        <tr>
                            <td>Payload:</td>
                            <td><textarea class="form-control form-control-sm" cols="10" rows="5"
                                          id="payload">
{
    "some": ["json","payload"]
}
                            </textarea></td>
                        </tr>
                        <tr>
                            <td></td>
                            <td><input class="btn btn-success" style="float:right;" type="submit"></td>
                        </tr>
                    </table>
                </form>
            </div>
        </div>
    </div>
{{end}}