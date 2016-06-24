var plus = document.getElementById("#plus");

plus.addEventListener("mousedown", function onCLick() {
	var field = document.getElementById("#text");
	console.log(field);
	document.body.append(field);
})