var bills = [
	{
		value: 1,
		numberOfBills: 20
	},
	{
		value: 5,
		numberOfBills: 20
	},
	{
		value: 10,
		numberOfBills: 20
	},
	{
		value: 20,
		numberOfBills: 20
	},
	{
		value: 50,
		numberOfBills: 20
	},
	{
		value: 100,
		numberOfBills: 20
	}
];

function getTotalAmount() {
	var total = 0;
	for (i = 0; i < bills.length; i++) {
		total += ( bills[i].value * bills[i].numberOfBills );
	}

	console.log('Cash Register contains: $' + total + '.');
}

var inputDenomination;
var inputNumber;

function choseDenomination() {
	inputDenomination = document.getElementById("input1").value;
}

function choseNumberOfBills() {
	inputNumber = document.getElementById("input2").value;
}


//fix this
//inputNumber is a string so the math is breaking
function add() {
	for (i = 0; i < bills.length; i++) {
		if (inputDenomination === bills[i].value) {
			bills[i].numberOfBills += inputNumber;
		}
	}
}