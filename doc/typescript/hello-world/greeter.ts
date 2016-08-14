function greeter(person) {
    return "Hello World! Hello, " + person;
}
var user = "Gyu-Ho Lee";
document.getElementsByClassName('greeter')[0].innerHTML = greeter(user);

// typed with TypeScript
function greeterTyped(person: string) {
    return "Hello World! Hello, " + person;
}
var userTyped = "Gyu-Ho Lee";
document.getElementsByClassName('greeterTyped')[0].innerHTML = greeterTyped(userTyped);

// interface with TypeScript
interface Person {
    firstName: string;
    lastName: string;
}
function greeterInterface(person: Person) {
    return "Hello, " + person.firstName + " " + person.lastName;
}
var userInterface = { firstName: "Gyu-Ho", lastName: "Lee" };
document.getElementsByClassName('greeterInterface')[0].innerHTML = greeterInterface(userInterface);

// class with TypeScript
class Student {
    fullName: string;
    constructor(public firstName, public middleInitial, public lastName) {
        this.fullName = firstName + " " + middleInitial + " " + lastName;
    }
}
var userClass = new Student("Gyu-Ho", "...", "Lee");
document.getElementsByClassName('greeterClass')[0].innerHTML = greeterInterface(userClass);
