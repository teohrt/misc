// user input for bill type
var inputDenomination;

// user input for number of those bills
var inputNumber;

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

function choseDenomination() {
	inputDenomination = document.getElementById("input1").value;
}

function choseNumberOfBills() {
	inputNumber = document.getElementById("input2").value;
}

function add() {
	for (i = 0; i < bills.length; i++) {
		if (parseInt(inputDenomination) == bills[i].value) {
			bills[i].numberOfBills += parseInt(inputNumber);
		}
	}
}

function subtract() {
	for (i = 0; i < bills.length; i++) {
		if (parseInt(inputDenomination) == bills[i].value) {
			bills[i].numberOfBills -= parseInt(inputNumber);
		}
	}
}