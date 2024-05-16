package examples

import (
	"context"
	"fmt"

	"github.com/gnolang/gno/gno.me/gno"
)

func CreateRemoteInstallerApp(vm gno.VM) error {
	appCode := fmt.Sprintf(remoteAppDefinition, "`"+remoteRenderContents+"`")
	return vm.Create(context.Background(), appCode, false)
}

const remoteAppDefinition = `
package remoteinstaller

func Render(_ string) string {
	return %s
}
`

const remoteRenderContents = `
<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Create App</title>
	<script>
		function submitForm() {
			var formData = {
				name: document.getElementById("name").value,
				address: document.getElementById("address").value
			};

			fetch('http://localhost:4591/system/create', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					'Access-Control-Allow-Origin': '*'
				},
				body: JSON.stringify(formData)
			})
				.then(response => {
					console.log(response);
				})
				.catch(error => {
					console.error('Error:', error);
				});
		}
	</script>
</head>

<body>
	<h2>Install Remote App</h2>
	<form id="myForm">
		<label for="address">Address:</label><br>
		<input type="text" id="address" name="address"><br><br>
		<label for="name">Name:</label><br>
		<input type="text" id="name" name="name"><br><br>
		<input type="button" value="Submit" onclick="submitForm()">
	</form>
</body>

</html>
`
