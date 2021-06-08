(()=>{
	const ingredients = document.getElementById("ingredients");
	const ingredientNamePattern = /[a-z]([a-z ]+[a-z])?/;
	let rows = [document.getElementById("ingredient-row")];

	setupInputs(rows[0]);

	function handleInput() {
		if (rows.every(isFull)) {
			addRow();
		} else {
			const emptyIndices = gatherEmptyRowIndices();
			// Leave first empty row intact
			emptyIndices.shift();
			// Filter out other empty rows
			if (emptyIndices.length > 0) {
				rows = rows.filter((row, idx) => {
					if (emptyIndices.length > 0 && idx === emptyIndices[0]) {
						emptyIndices.shift();
						row.remove();
						return false;
					}
					return true;
				});
			}
		}
	}
	
	function isEmpty(row) {
		for (let i = 0; i < row.children.length; ++i) {
			if (row.children[i] instanceof HTMLInputElement && row.children[i].value.trim() !== '') {
				return false;
			}
		}
		return true;
	}

	function isFull(row) {
		for (let i = 0; i < row.children.length; ++i) {
			if (row.children[i] instanceof HTMLInputElement && row.children[i].value.trim() === '') {
				return false;
			}
		}
		return true;
	}

	function gatherEmptyRowIndices() {
		const result = [];
		rows.forEach((row, idx) => {
			if (isEmpty(row)) {
				result.push(idx);
			}
		});
		return result;
	}

	function setupInputs(row) {
		for (let i = 0; i < row.children.length; ++i) {
			if (row.children[i] instanceof HTMLInputElement) {
				row.children[i].value = '';
				row.children[i].addEventListener('input', handleInput);
			}
		}
	}

	/** Append a new ingredient input row. */
	function addRow() {
		const row = rows[0].cloneNode(true);
		setupInputs(row)
		rows.push(row);
		ingredients.appendChild(row);
	}
})();
