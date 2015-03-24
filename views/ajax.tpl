{{define "ajax"}}
{{range .Articles}}
  <div class="index-post">
    <h3>{{.Title}}</h3>
    <div class="index-post-meta">{{.Date}} by {{.Author}} 点击：{{.Clink}}</div>
    <blockquote>
    <p>{{str2html .Lead}}</p>
    <a type="button" data-toggle="modal" data-target="#myModal">[详情]</a>
    </blockquote>
    <div class="modal fade bs-example-modal-lg" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h3 class="modal-title" id="myModalLabel">{{.Title}}</h3>
            <div class="index-post-meta">{{.Date}} by {{.Author}} 点击：{{.Clink}}</div>
          </div>
          <div class="modal-body">
            {{str2html .Content}}
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关 闭</button>
          </div>
        </div>
      </div>
    </div>
    <hr>
  </div><!-- /.index-post -->
{{end}}
<script type="text/javascript">
function loadXMLDoc()
{
  var xmlhttp;
  if (window.XMLHttpRequest)
    {// code for IE7+, Firefox, Chrome, Opera, Safari
    xmlhttp=new XMLHttpRequest();
    }
  else
    {// code for IE6, IE5
    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }
    xmlhttp.onreadystatechange=function()
      {
      if (xmlhttp.readyState==4 && xmlhttp.status==200)
        {
        document.getElementById("myDiv").innerHTML=xmlhttp.responseText;
        }
      }
    xmlhttp.open("GET","/getajaxarticles?Id=181&Pag=37",true);
    xmlhttp.send();
  alert("Ajax!");
}
</script>
<div id="myDiv"></div>
{{end}}
