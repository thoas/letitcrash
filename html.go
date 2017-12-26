package letitgo

const debugTpl = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
    <meta name="robots" content="NONE,NOARCHIVE">
    <title>Error at {{ .request.URL.Path }}</title>
    <style type="text/css">
        html * {
            padding: 0;
            margin: 0;
        }
        
        body * {
            padding: 10px 20px;
        }
        
        body * * {
            padding: 0;
        }
        
        body {
            font: small sans-serif;
            background-color: #fff;
            color: #000;
        }
        
        body>div {
            border-bottom: 1px solid #ddd;
        }
        
        h1 {
            font-weight: normal;
        }
        
        h2 {
            margin-bottom: .8em;
        }
        
        h2 span {
            font-size: 80%;
            color: #666;
            font-weight: normal;
        }
        
        h3 {
            margin: 1em 0 .5em 0;
        }
        
        h4 {
            margin: 0 0 .5em 0;
            font-weight: normal;
        }
        
        code,
        pre {
            font-size: 100%;
            white-space: pre-wrap;
        }
        
        table {
            border: 1px solid #ccc;
            border-collapse: collapse;
            width: 100%;
            background: white;
        }
        
        tbody td,
        tbody th {
            vertical-align: top;
            padding: 2px 3px;
        }
        
        thead th {
            padding: 1px 6px 1px 3px;
            background: #fefefe;
            text-align: left;
            font-weight: normal;
            font-size: 11px;
            border: 1px solid #ddd;
        }
        
        tbody th {
            width: 12em;
            text-align: right;
            color: #666;
            padding-right: .5em;
        }
        
        table.vars {
            margin: 5px 0 2px 40px;
        }
        
        table.vars td,
        table.req td {
            font-family: monospace;
        }
        
        table td.code {
            width: 100%;
        }
        
        table td.code pre {
            overflow: hidden;
        }
        
        table.source th {
            color: #666;
        }
        
        table.source td {
            font-family: monospace;
            white-space: pre;
            border-bottom: 1px solid #eee;
        }
        
        ul.traceback {
            list-style-type: none;
            color: #222;
        }
        
        ul.traceback li.frame {
            padding-bottom: 1em;
            color: #666;
        }
        
        ul.traceback li.user {
            background-color: #e0e0e0;
            color: #000
        }
        
        div.context {
            padding: 10px 0;
            overflow: hidden;
        }
        
        div.context ol {
            padding-left: 30px;
            margin: 0 10px;
            list-style-position: inside;
        }
        
        div.context ol li {
            font-family: monospace;
            white-space: pre;
            color: #777;
            cursor: pointer;
            padding-left: 2px;
        }
        
        div.context ol li pre {
            display: inline;
        }
        
        div.context ol.context-line li {
            color: #505050;
            background-color: #dfdfdf;
            padding: 3px 2px;
        }
        
        div.context ol.context-line li span {
            position: absolute;
            right: 32px;
        }
        
        .user div.context ol.context-line li {
            background-color: #bbb;
            color: #000;
        }
        
        .user div.context ol li {
            color: #666;
        }
        
        div.commands {
            margin-left: 40px;
        }
        
        div.commands a {
            color: #555;
            text-decoration: none;
        }
        
        .user div.commands a {
            color: black;
        }
        
        #summary {
            background: #E0EBF5;
        }
        
        #summary h2 {
            font-weight: normal;
            color: #666;
        }
        
        #explanation {
            background: #eee;
        }
        
        #template,
        #template-not-exist {
            background: #f6f6f6;
        }
        
        #template-not-exist ul {
            margin: 0 0 10px 20px;
        }
        
        #template-not-exist .postmortem-section {
            margin-bottom: 3px;
        }
        
        #unicode-hint {
            background: #eee;
        }
        
        #traceback {
            background: #eee;
        }
        
        #requestinfo {
            background: #f6f6f6;
            padding-left: 120px;
        }

        #serverinfo {
            background: #f6f6f6;
            padding-left: 200px;
        }
        
        #summary table {
            border: none;
            background: transparent;
        }
        
        #requestinfo h2,
        #requestinfo h3 {
            position: relative;
            margin-left: -100px;
        }

        #serverinfo h2,
		#serverinfo h3 {
            position: relative;
            margin-left: -180px;
        }
        
        #requestinfo h3, #serverinfo h3 {
            margin-bottom: -1em;
        }
        
        .error {
            background: #ffc;
        }
        
        .specific {
            color: #cc3300;
            font-weight: bold;
        }
        
        h2 span.commands {
            font-size: .7em;
        }
        
        span.commands a:link {
            color: #5E5694;
        }
        
        pre.exception_value {
            font-family: sans-serif;
            color: #666;
            font-size: 1.5em;
            margin: 10px 0 10px 0;
        }
        
        .append-bottom {
            margin-bottom: 10px;
        }
    </style>
    <script type="text/javascript">
        function hideAll(elems) {
            for (var e = 0; e < elems.length; e++) {
                elems[e].style.display = 'none';
            }
        }
        window.onload = function() {
            hideAll(document.querySelectorAll('table.vars'));
            hideAll(document.querySelectorAll('ol.pre-context'));
            hideAll(document.querySelectorAll('ol.post-context'));
            hideAll(document.querySelectorAll('div.pastebin'));
        }

        function toggle() {
            for (var i = 0; i < arguments.length; i++) {
                var e = document.getElementById(arguments[i]);
                if (e) {
                    e.style.display = e.style.display == 'none' ? 'block' : 'none';
                }
            }
            return false;
        }

        function varToggle(link, id) {
            toggle('v' + id);
            var s = link.getElementsByTagName('span')[0];
            var uarr = String.fromCharCode(0x25b6);
            var darr = String.fromCharCode(0x25bc);
            s.textContent = s.textContent == uarr ? darr : uarr;
            return false;
        }

        function switchPastebinFriendly(link) {
            s1 = "Switch to copy-and-paste view";
            s2 = "Switch back to interactive view";
            link.textContent = link.textContent.trim() == s1 ? s2 : s1;
            toggle('browserTraceback', 'pastebinTraceback');
            return false;
        }
    </script>
    <style></style>
</head>

<body>
    <div id="summary">
        <h1>Error at {{ .request.URL.Path }}</h1>
		<pre class="exception_value">{{ .error }}</pre>
        <table class="meta">
            <tbody>
                <tr>
                    <th>Request Method:</th>
                    <td>{{ .request.Method }}</td>
                </tr>
                <tr>
                    <th>Request URL:</th>
                    <td>http://{{ .request.Host }}{{ .request.URL }}</td>
                </tr>
                <tr>
                    <th>Error Type:</th>
                    <td>{{ .errorType }}</td>
                </tr>
                <tr>
                    <th>Error Value:</th>
                    <td><pre>{{ .error }}</pre></td>
                </tr>
                <tr>
                    <th>Exception Location:</th>
                    <td>src/forj/web/frontend/views.py in healthcheck, line 69</td>
                </tr>
                <tr>
                    <th>Go Version:</th>
                    <td>{{ .version }}</td>
                </tr>
                <tr>
                    <th>Server time:</th>
                    <td>{{ .now }}</td>
                </tr>
            </tbody>
        </table>
    </div>
    <div id="traceback">
        <h2>Traceback <span class="commands"><a href="#" onclick="return switchPastebinFriendly(this);">Switch to copy-and-paste view</a></span></h2>
        <div id="browserTraceback" style="display: block;">
            <ul class="traceback">
                <li class="frame django"><code>/Users/thoas/Sites/Python/forj/.env/lib/python3.6/site-packages/django/core/handlers/exception.py</code> in <code>inner</code>
                    <div class="context" id="c4467487432">
                        <ol start="28" class="pre-context" id="pre4467487432" style="display: none;">
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>    This decorator is automatically applied to all middleware to ensure that</pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>    no middleware leaks an exception and that the next middleware in the stack</pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>    can rely on getting a response instead of an exception.</pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>    """</pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>    @wraps(get_response)</pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>    def inner(request):</pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>        try:</pre></li>
                        </ol>
                        <ol start="35" class="context-line">
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>            response = get_response(request)</pre><span>...</span></li>
                        </ol>
                        <ol start="36" class="post-context" id="post4467487432" style="display: none;">
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>        except Exception as exc:</pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>            response = response_for_exception(request, exc)</pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>        return response</pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre>    return inner</pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre></pre></li>
                            <li onclick="toggle('pre4467487432', 'post4467487432')"><pre></pre></li>
                        </ol>
                    </div>
                    <div class="commands"><a href="#" onclick="return varToggle(this, '4467487432')"><span>▶</span> Local vars</a></div>
                    <table class="vars" id="v4467487432" style="display: none;">
                        <thead>
                            <tr>
                                <th>Variable</th>
                                <th>Value</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>exc</td>
                                <td class="code"><pre>Exception('foo',)</pre></td>
                            </tr>
                            <tr>
                                <td>get_response</td>
                                <td class="code"><pre>&lt;bound method BaseHandler._get_response of &lt;django.core.handlers.wsgi.WSGIHandler object at 0x10a41e588&gt;&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>request</td>
                                <td class="code"><pre>&lt;WSGIRequest: GET '/healthcheck/'&gt;</pre></td>
                            </tr>
                        </tbody>
                    </table>
                </li>
                <li class="frame django"><code>/Users/thoas/Sites/Python/forj/.env/lib/python3.6/site-packages/django/core/handlers/base.py</code> in <code>_get_response</code>
                    <div class="context" id="c4467487368">
                        <ol start="121" class="pre-context" id="pre4467487368" style="display: none;">
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>                break</pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre></pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>        if response is None:</pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>            wrapped_callback = self.make_view_atomic(callback)</pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>            try:</pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>                response = wrapped_callback(request, *callback_args, **callback_kwargs)</pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>            except Exception as e:</pre></li>
                        </ol>
                        <ol start="128" class="context-line">
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>                response = self.process_exception_by_middleware(e, request)</pre><span>...</span></li>
                        </ol>
                        <ol start="129" class="post-context" id="post4467487368" style="display: none;">
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre></pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>        # Complain if the view returned None (a common error).</pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>        if response is None:</pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>            if isinstance(callback, types.FunctionType):    # FBV</pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>                view_name = callback.__name__</pre></li>
                            <li onclick="toggle('pre4467487368', 'post4467487368')"><pre>            else:                                           # CBV</pre></li>
                        </ol>
                    </div>
                    <div class="commands"><a href="#" onclick="return varToggle(this, '4467487368')"><span>▶</span> Local vars</a></div>
                    <table class="vars" id="v4467487368" style="display: none;">
                        <thead>
                            <tr>
                                <th>Variable</th>
                                <th>Value</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>callback</td>
                                <td class="code"><pre>&lt;function healthcheck at 0x10a119bf8&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>callback_args</td>
                                <td class="code"><pre>()</pre></td>
                            </tr>
                            <tr>
                                <td>callback_kwargs</td>
                                <td class="code"><pre>{}</pre></td>
                            </tr>
                            <tr>
                                <td>middleware_method</td>
                                <td class="code"><pre>&lt;bound method CsrfViewMiddleware.process_view of &lt;django.middleware.csrf.CsrfViewMiddleware object at 0x10a431160&gt;&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>request</td>
                                <td class="code"><pre>&lt;WSGIRequest: GET '/healthcheck/'&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>resolver</td>
                                <td class="code"><pre>&lt;URLResolver 'forj.web.frontend.urls' (None:None) '^/'&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>resolver_match</td>
                                <td class="code"><pre>ResolverMatch(func=forj.web.frontend.views.healthcheck, args=(), kwargs={}, url_name=healthcheck, app_names=[], namespaces=[])</pre></td>
                            </tr>
                            <tr>
                                <td>response</td>
                                <td class="code"><pre>None</pre></td>
                            </tr>
                            <tr>
                                <td>self</td>
                                <td class="code"><pre>&lt;django.core.handlers.wsgi.WSGIHandler object at 0x10a41e588&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>urlconf</td>
                                <td class="code"><pre>'forj.web.frontend.urls'</pre></td>
                            </tr>
                            <tr>
                                <td>wrapped_callback</td>
                                <td class="code"><pre>&lt;function healthcheck at 0x10a119bf8&gt;</pre></td>
                            </tr>
                        </tbody>
                    </table>
                </li>
                <li class="frame django"><code>/Users/thoas/Sites/Python/forj/.env/lib/python3.6/site-packages/django/core/handlers/base.py</code> in <code>_get_response</code>
                    <div class="context" id="c4467487304">
                        <ol start="119" class="pre-context" id="pre4467487304" style="display: none;">
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>            response = middleware_method(request, callback, callback_args, callback_kwargs)</pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>            if response:</pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>                break</pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre></pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>        if response is None:</pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>            wrapped_callback = self.make_view_atomic(callback)</pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>            try:</pre></li>
                        </ol>
                        <ol start="126" class="context-line">
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>                response = wrapped_callback(request, *callback_args, **callback_kwargs)</pre><span>...</span></li>
                        </ol>
                        <ol start="127" class="post-context" id="post4467487304" style="display: none;">
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>            except Exception as e:</pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>                response = self.process_exception_by_middleware(e, request)</pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre></pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>        # Complain if the view returned None (a common error).</pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>        if response is None:</pre></li>
                            <li onclick="toggle('pre4467487304', 'post4467487304')"><pre>            if isinstance(callback, types.FunctionType):    # FBV</pre></li>
                        </ol>
                    </div>
                    <div class="commands"><a href="#" onclick="return varToggle(this, '4467487304')"><span>▶</span> Local vars</a></div>
                    <table class="vars" id="v4467487304" style="display: none;">
                        <thead>
                            <tr>
                                <th>Variable</th>
                                <th>Value</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>callback</td>
                                <td class="code"><pre>&lt;function healthcheck at 0x10a119bf8&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>callback_args</td>
                                <td class="code"><pre>()</pre></td>
                            </tr>
                            <tr>
                                <td>callback_kwargs</td>
                                <td class="code"><pre>{}</pre></td>
                            </tr>
                            <tr>
                                <td>middleware_method</td>
                                <td class="code"><pre>&lt;bound method CsrfViewMiddleware.process_view of &lt;django.middleware.csrf.CsrfViewMiddleware object at 0x10a431160&gt;&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>request</td>
                                <td class="code"><pre>&lt;WSGIRequest: GET '/healthcheck/'&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>resolver</td>
                                <td class="code"><pre>&lt;URLResolver 'forj.web.frontend.urls' (None:None) '^/'&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>resolver_match</td>
                                <td class="code"><pre>ResolverMatch(func=forj.web.frontend.views.healthcheck, args=(), kwargs={}, url_name=healthcheck, app_names=[], namespaces=[])</pre></td>
                            </tr>
                            <tr>
                                <td>response</td>
                                <td class="code"><pre>None</pre></td>
                            </tr>
                            <tr>
                                <td>self</td>
                                <td class="code"><pre>&lt;django.core.handlers.wsgi.WSGIHandler object at 0x10a41e588&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>urlconf</td>
                                <td class="code"><pre>'forj.web.frontend.urls'</pre></td>
                            </tr>
                            <tr>
                                <td>wrapped_callback</td>
                                <td class="code"><pre>&lt;function healthcheck at 0x10a119bf8&gt;</pre></td>
                            </tr>
                        </tbody>
                    </table>
                </li>
                <li class="frame django"><code>/Users/thoas/Sites/Python/forj/.env/lib/python3.6/site-packages/django/views/decorators/csrf.py</code> in <code>wrapped_view</code>
                    <div class="context" id="c4467486792">
                        <ol start="47" class="pre-context" id="pre4467486792" style="display: none;">
                            <li onclick="toggle('pre4467486792', 'post4467486792')"><pre></pre></li>
                            <li onclick="toggle('pre4467486792', 'post4467486792')"><pre></pre></li>
                            <li onclick="toggle('pre4467486792', 'post4467486792')"><pre>def csrf_exempt(view_func):</pre></li>
                            <li onclick="toggle('pre4467486792', 'post4467486792')"><pre>    """Mark a view function as being exempt from the CSRF view protection."""</pre></li>
                            <li onclick="toggle('pre4467486792', 'post4467486792')"><pre>    # view_func.csrf_exempt = True would also work, but decorators are nicer</pre></li>
                            <li onclick="toggle('pre4467486792', 'post4467486792')"><pre>    # if they don't have side effects, so return a new function.</pre></li>
                            <li onclick="toggle('pre4467486792', 'post4467486792')"><pre>    def wrapped_view(*args, **kwargs):</pre></li>
                        </ol>
                        <ol start="54" class="context-line">
                            <li onclick="toggle('pre4467486792', 'post4467486792')"><pre>        return view_func(*args, **kwargs)</pre><span>...</span></li>
                        </ol>
                        <ol start="55" class="post-context" id="post4467486792" style="display: none;">
                            <li onclick="toggle('pre4467486792', 'post4467486792')"><pre>    wrapped_view.csrf_exempt = True</pre></li>
                            <li onclick="toggle('pre4467486792', 'post4467486792')"><pre>    return wraps(view_func)(wrapped_view)</pre></li>
                        </ol>
                    </div>
                    <div class="commands"><a href="#" onclick="return varToggle(this, '4467486792')"><span>▶</span> Local vars</a></div>
                    <table class="vars" id="v4467486792" style="display: none;">
                        <thead>
                            <tr>
                                <th>Variable</th>
                                <th>Value</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>args</td>
                                <td class="code"><pre>(&lt;WSGIRequest: GET '/healthcheck/'&gt;,)</pre></td>
                            </tr>
                            <tr>
                                <td>kwargs</td>
                                <td class="code"><pre>{}</pre></td>
                            </tr>
                            <tr>
                                <td>view_func</td>
                                <td class="code"><pre>&lt;function healthcheck at 0x10a041b70&gt;</pre></td>
                            </tr>
                        </tbody>
                    </table>
                </li>
                <li class="frame user"><code>src/forj/web/frontend/views.py</code> in <code>healthcheck</code>
                    <div class="context" id="c4467487240">
                        <ol start="62" class="pre-context" id="pre4467487240" style="display: block;">
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre>    results['sha'] = release_tag</pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre>    results['tznow'] = timezone.now()</pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre>    results['now'] = datetime.now()</pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre>    results['version'] = version</pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre>    results['uptime'] = uptime</pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre>    results['uptime_since'] = naturaltime(uptime)</pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre></pre></li>
                        </ol>
                        <ol start="69" class="context-line">
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre>    raise Exception('foo')</pre><span>...</span></li>
                        </ol>
                        <ol start="70" class="post-context" id="post4467487240" style="display: block;">
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre></pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre>    return JsonResponse(results)</pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre></pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre></pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre>def home(request, template_name='forj/home.html', **extra_context):</pre></li>
                            <li onclick="toggle('pre4467487240', 'post4467487240')"><pre>    cart = Cart.from_request(request)</pre></li>
                        </ol>
                    </div>
                    <div class="commands"><a href="#" onclick="return varToggle(this, '4467487240')"><span>▼</span> Local vars</a></div>
                    <table class="vars" id="v4467487240" style="display: block;">
                        <thead>
                            <tr>
                                <th>Variable</th>
                                <th>Value</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>release_tag</td>
                                <td class="code"><pre>None</pre></td>
                            </tr>
                            <tr>
                                <td>request</td>
                                <td class="code"><pre>&lt;WSGIRequest: GET '/healthcheck/'&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>results</td>
                                <td class="code"><pre>{'GET': &lt;QueryDict: {}&gt;,
 'HOST': 'www.local.forj.shop:8181',
 'HTTP_ACCEPT_LANGUAGE': 'en,en-GB;q=0.9,en-US;q=0.8,fr;q=0.7',
 'HTTP_HOST': 'www.local.forj.shop:8181',
 'HTTP_X_FORWARDED_FOR': None,
 'HTTP_X_FORWARDED_PROTO': None,
 'HTTP_X_FORWARDED_PROTOCOL': None,
 'IS_AJAX': False,
 'IS_SECURE': False,
 'POST': &lt;QueryDict: {}&gt;,
 'QUERY_STRING': '',
 'REMOTE_ADDR': '127.0.0.1',
 'USER': 'AnonymousUser',
 'X-Real-Ip': None,
 'now': datetime.datetime(2017, 12, 25, 18, 59, 43, 547962),
 'sha': None,
 'tznow': datetime.datetime(2017, 12, 25, 17, 59, 43, 547953, tzinfo=&lt;UTC&gt;),
 'uptime': datetime.datetime(2017, 12, 25, 18, 59, 42, 226639),
 'uptime_since': 'a second ago',
 'version': '0.1.0'}</pre></td>
                            </tr>
                            <tr>
                                <td>uptime</td>
                                <td class="code"><pre>datetime.datetime(2017, 12, 25, 18, 59, 42, 226639)</pre></td>
                            </tr>
                            <tr>
                                <td>user</td>
                                <td class="code"><pre>&lt;SimpleLazyObject: &lt;django.contrib.auth.models.AnonymousUser object at 0x10a482588&gt;&gt;</pre></td>
                            </tr>
                            <tr>
                                <td>version</td>
                                <td class="code"><pre>'0.1.0'</pre></td>
                            </tr>
                        </tbody>
                    </table>
                </li>
            </ul>
        </div>
        <form action="http://dpaste.com/" name="pasteform" id="pasteform" method="post">
            <div id="pastebinTraceback" class="pastebin" style="display: none;">
                <input type="hidden" name="language" value="PythonConsole">
                <input type="hidden" name="title" value="Exception at /healthcheck/">
                <input type="hidden" name="source" value="Django Dpaste Agent">
                <input type="hidden" name="poster" value="Django">
                <textarea name="content" id="traceback_area" cols="140" rows="25">Environment: Request Method: GET Request URL: http://www.local.forj.shop:8181/healthcheck/ Django Version: 2.0 Python Version: 3.6.3 Installed Applications: ['django.contrib.admin', 'django.contrib.auth', 'django.contrib.contenttypes', 'django.contrib.messages', 'django.contrib.staticfiles', 'forj', 'django_extensions', 'easy_thumbnails', 'django_jinja', 'django_jinja.contrib._easy_thumbnails', 'django_jinja.contrib._humanize', 'django_hosts'] Installed Middleware: ['django_hosts.middleware.HostsRequestMiddleware', 'django.middleware.security.SecurityMiddleware', 'django.contrib.sessions.middleware.SessionMiddleware', 'django.middleware.common.CommonMiddleware', 'django.middleware.csrf.CsrfViewMiddleware', 'django.contrib.auth.middleware.AuthenticationMiddleware', 'django.contrib.messages.middleware.MessageMiddleware', 'django.middleware.clickjacking.XFrameOptionsMiddleware', 'django_hosts.middleware.HostsResponseMiddleware', 'forj.middleware.MinifyHTMLMiddleware', 'forj.middleware.SetRemoteAddrFromForwardedFor'] Traceback: File "/Users/thoas/Sites/Python/forj/.env/lib/python3.6/site-packages/django/core/handlers/exception.py" in inner 35. response = get_response(request) File "/Users/thoas/Sites/Python/forj/.env/lib/python3.6/site-packages/django/core/handlers/base.py" in _get_response 128. response = self.process_exception_by_middleware(e, request) File "/Users/thoas/Sites/Python/forj/.env/lib/python3.6/site-packages/django/core/handlers/base.py" in _get_response 126. response = wrapped_callback(request, *callback_args, **callback_kwargs) File "/Users/thoas/Sites/Python/forj/.env/lib/python3.6/site-packages/django/views/decorators/csrf.py" in wrapped_view 54. return view_func(*args, **kwargs) File "src/forj/web/frontend/views.py" in healthcheck 69. raise Exception('foo') Exception Type: Exception at /healthcheck/ Exception Value: foo
                </textarea>
                <br>
                <br>
                <input type="submit" value="Share this traceback on a public website">
            </div>
        </form>
    </div>
    <div id="requestinfo">
        <h2>Request information</h2>
        <h3 id="header-info">HEADER</h3>
		{{ if .headers }}
        <table class="req">
            <thead>
                <tr>
                    <th>Variable</th>
                    <th>Value</th>
                </tr>
            </thead>
            <tbody>
			{{ range $key, $value := .headers }}
			<tr>
			<td>{{ $key }}</td>
			<td class="code"><pre>{{ $value }}</pre></td>
			</tr>
			{{ end }}
            </tbody>
        </table>
		{{ else }}
        <p>No HEADER data</p>
		{{ end }}
        <h3 id="get-info">GET</h3>
		{{ if .queryString }}
        <table class="req">
            <thead>
                <tr>
                    <th>Variable</th>
                    <th>Value</th>
                </tr>
            </thead>
            <tbody>
			{{ range $key, $value := .queryString }}
			<tr>
			<td>{{ $key }}</td>
			<td class="code"><pre>{{ $value }}</pre></td>
			</tr>
			{{ end }}
            </tbody>
        </table>
		{{ else }}
        <p>No GET data</p>
		{{ end }}
        <h3 id="post-info">POST</h3>
		{{ if .data }}
		{{ else }}
        <p>No POST data</p>
		{{ end }}
        <h3 id="files-info">FILES</h3>
        <p>No FILES data</p>
        <h3 id="cookie-info">COOKIES</h3>
		{{ if .cookies }}
        <table class="req">
            <thead>
                <tr>
                    <th>Variable</th>
                    <th>Value</th>
                </tr>
            </thead>
            <tbody>
				{{ range $key, $value := .cookies }}
                <tr>
                    <td>{{ $value.Name }}</td>
                    <td class="code"><pre>{{ $value.Value }}</pre></td>
                </tr>
				{{ end }}
            </tbody>
        </table>
		{{ else }}
        <p>No COOKIES data</p>
		{{ end }}
        <h3 id="environment-info">ENV</h3>
		{{ if .environments }}
        <table class="req">
            <thead>
                <tr>
                    <th>Variable</th>
                    <th>Value</th>
                </tr>
            </thead>
            <tbody>
				{{ range $key, $value := .environments }}
                <tr>
                    <td>{{ $key }}</td>
                    <td class="code"><pre>{{ $value }}</pre></td>
                </tr>
				{{ end }}
            </tbody>
        </table>
		{{ else }}
        <p>No ENVIRONMENTS data</p>
		{{ end }}
	</div>
    <div id="serverinfo">
        <h2>Server information</h2>
		{{ range $key, $value := .serverInformation }}
        <h3>{{ $key }}</h3>
        <p>{{ $value }}</p>
		{{ end }}
	</div>
</body>
</html>`
