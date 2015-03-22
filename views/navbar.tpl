{{define "navbar"}}
<nav class="navbar navbar-inverse navbar-fixed-top">
  <div class="container">
    <div id="navbar" class="navbar-collapse collapse">
      <ul class="nav navbar-nav">
        <li {{if .Nav1.Active}}class="active"{{end}}><a href="{{.Nav1.Href}}">{{.Nav1.Name}}</a></li>
      	{{range .Nav2}}
        <li {{if .Active}}class="active"{{end}}><a href="{{.Href}}">{{.Name}}</a></li>
        {{end}}
        <li class="dropdown">
        	<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">更 多 <span class="caret"></span></a>
          <ul class="dropdown-menu" role="menu">
				    {{range .Nav3}}
        		<li {{if .Active}}class="active"{{end}}><a href="{{.Href}}">{{.Name}}</a></li>
        		{{end}}
			     </ul>
        </li>
      </ul>
      <form class="navbar-form navbar-right">
        <div class="form-group">
          <input type="text" placeholder="登录后搜索" class="form-control">
        </div>
        <div class="form-group">
          <input type="password" placeholder="请输入密码" class="form-control">
        </div>
        <button type="submit" class="btn btn-success">确 定</button>
      </form>
    </div>
  </div>
</nav>
{{end}}