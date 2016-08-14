function greeter1(person) {
    return "Hello World! Hello, " + person;
}
var user1 = "Gyu-Ho Lee 1";
document.body.innerHTML = greeter1(user1);
// typed with TypeScript
function greeterTyped(person) {
    return "Hello World 2! Hello, " + person;
}
var userTyped = "Gyu-Ho Lee";
document.getElementsByClassName('greeterTyped')[0].innerHTML = greeterTyped(userTyped);
function greeterInterface(person) {
    return "Hello, " + person.firstName + " " + person.lastName;
}
var userInterface = { firstName: "Gyu-Ho", lastName: "Lee" };
document.getElementsByClassName('greeterInterface')[0].innerHTML = greeterInterface(userInterface);
// class with TypeScript
var Student = (function () {
    function Student(firstName, middleInitial, lastName) {
        this.firstName = firstName;
        this.middleInitial = middleInitial;
        this.lastName = lastName;
        this.fullName = firstName + " " + middleInitial + " " + lastName;
    }
    return Student;
}());
var userClass = new Student("Gyu-Ho", "...", "Lee");
document.getElementsByClassName('greeterClass')[0].innerHTML = greeterInterface(userClass);
