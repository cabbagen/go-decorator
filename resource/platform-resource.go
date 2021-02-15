package resource

var MobileTpl = `
<html>
	<head>
		<meta charset="utf8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
    	<meta name="viewport" content="width=device-width,initial-scale=1.0">
		<title>{{.Title}}</title>
		<script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js" type="text/javascript"></script>
		<script src="http://localhost:7001/static/library/topic-component.min.js" type="text/javascript"></script>
	</head>
	<body>
		<div id="root"></div>
	</body>
	<script type="text/javascript">
		window._data = {{.Data}}
	</script>
	<script type="text/javascript">
		console.log(window._data);
	</script>
</html>
`
