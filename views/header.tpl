{{define "header"}}
	<!DOCTYPE html>
	<html lang="zh-cn">
	  <head>
	    <meta charset="utf-8">
	    <link rel="icon" href="static/img/favicon.ico">
	    <!-- Bootstrap core CSS -->
	    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
	    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-datetimepicker.min.css">
	    <!-- Custom styles for this template -->
	    <link href="/static/css/index.css" rel="stylesheet">
	    <!-- jquery 1.11.2-->
	    <script type="text/javascript" src="/static/js/jquery.min.js"></script>
	    <!-- bootstrap js -->
	    <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>

	    <script type="text/javascript" src="/static/js/bootstrap-datetimepicker.min.js"></script>
	    <script type="text/javascript" src="/static/js/locales/bootstrap-datetimepicker.zh-CN.js" charset="UTF-8"></script>
	    <title>价格在线 - {{.Title}}</title>
	  </head>
	  <body>
	  	<div class="index-masthead">
	  	  <div class="container">
	  	    {{template "navbar" .}}
	  	  </div>
	  	</div>
{{end}}