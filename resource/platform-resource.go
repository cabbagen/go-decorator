package resource

var MobileTpl = `
<html>
	<head>
		<meta charset="utf8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
    	<meta name="viewport" content="width=device-width,initial-scale=1.0">
		<title>{{.Title}}</title>
		<script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js" type="text/javascript"></script>
		<script src="http://39.106.83.126/library/topic-component.min.js" type="text/javascript"></script>
	</head>
	<body>
		<div id="root">
			<div id="root">
				<div v-for="(item, index) in data" :key="index">
					<component :is="'tc-' + item.type" v-bind="item" />
				</div>
			</div>
		</div>
	</body>
	<script type="text/javascript">
		window._data = {{.Data}}
	</script>
	<script type="text/javascript">
		var app = new Vue({
		    el: '#root',
		    data: {
				data: window._data.map(item => Object.assign({}, JSON.parse(item.content), { id: item.id, type: item.type })),
		    },
		});
	</script>
</html>
`
