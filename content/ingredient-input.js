addIngredient = (()=>{
	// TODO: switch to a template field
	const ingredientField = document.getElementById("ingredient-field");
	const ingredientSection = ingredientField.parentElement;

	return function addIngredient() {
		const field = ingredientField.cloneNode(true);
		ingredientSection.appendChild(field);
	}
})();
