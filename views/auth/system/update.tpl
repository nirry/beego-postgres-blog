<link href="/static/auth/js/validate/validate.css" rel="stylesheet">
<link rel="stylesheet" type="text/css" href="/static/auth/js/tag/jquery.tagsinput.css">
<link rel="stylesheet" href="/static/auth/css/editormd.min.css" />
<link rel="stylesheet" type="text/css" href="/static/auth/css/jquery-ui.min.css">

<div class="content-wrap">
    <div class="row">
        <div class="col-sm-12">
            <div class="nest" id="validationClose">
                <div class="title-alt">
                    <h6>
                        新建</h6>
                    <div class="titleClose">
                        <a class="gone" href="#validationClose">
                            <span class="entypo-cancel"></span>
                        </a>
                    </div>
                    <div class="titleToggle">
                        <a class="nav-toggle-alt" href="#validation">
                            <span class="entypo-up-open"></span>
                        </a>
                    </div>

                </div>

                <div class="body-nest" id="validation">
                    <div class="form_center">

                        <form action="/console/system/1" method="post" id="systemCreate" class="form-horizontal systemForm">
                        {{ .xsrfdata }}
                            <fieldset>
                                <input type="hidden" value="PUT" name="_method">
                                <div class="control-group">
                                    <label class="control-label" for="name">网站标题</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.Title}}" class="form-control" name="title" id="name">
                                    </div>
                                </div>
                                <div class="control-group">
                                    <label class="control-label" for="displayName">标题SEO名</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.STitle}}" class="form-control" name="s_title" id="displayName">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">网站描述</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.Description}}" class="form-control" name="description" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">网站SEO关键词</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.SeoKey}}" class="form-control" name="seo_key" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">网站SEO描述</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.SeoDes}}" class="form-control" name="seo_des" id="description">
                                    </div>
                                </div>
                                <div class="control-group">
                                    <label class="control-label" for="comment_type">评论类型</label>
                                    <div class="controls">
                                        <select class="form-control" name="comment_type"  id="comment_type">
                                        {{range $index,$item := .comment }}
                                            <option value="{{$index}}" {{if eq $index $.system.CommentType }} selected {{end}}>{{$item}}</option>
                                        {{end}}
                                        </select>
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">GithubClientId</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.GithubClientId}}" class="form-control" name="github_client_id" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">GithubClientSecret</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.GithubClientSecret}}" class="form-control" name="github_client_secret" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">GithubName</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.GithubName}}" class="form-control" name="github_name" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">GithubRepo</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.GithubRepo}}" class="form-control" name="github_repo" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">CyAppId</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.CyAppId}}" class="form-control" name="cy_app_id" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">CyAppKey</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.CyAppKey}}" class="form-control" name="cy_app_key" id="description">
                                    </div>
                                </div>
                                <div class="control-group">
                                    <label class="control-label" for="comment_type">CDN类型</label>
                                    <div class="controls">
                                        <select class="form-control" name="cdn_type"  id="comment_type">
                                        {{range $index,$item := .cdn }}
                                            <option value="{{$index}}" {{if eq $index $.system.CdnType }} selected {{end}}>{{$item}}</option>
                                        {{end}}
                                        </select>
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">CDN地址</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.CdnUrl}}" class="form-control" name="cdn_url" id="description">
                                    </div>
                                </div>
                                <div class="control-group" >
                                    <label class="control-label" for="description">网站备案号</label>
                                    <div  class="controls">
                                        <input type="text" value="{{.system.RecordNumber}}" class="form-control" name="record_number" id="description">
                                    </div>
                                </div>
                                <div class="form-actions" style="margin:20px 0 0 0;">
                                    <button type="submit" class="btn btn-primary">Submit</button>
                                    <button type="reset" class="btn">Cancel</button>
                                </div>
                            </fieldset>
                        </form>

                    </div>
                </div>

            </div>
        </div>
    </div>
</div>




</div>