package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("1. Create a program that displays the message", "Hello world", "on the screen.")

	fmt.Println("Hello World")

	fmt.Println("2. Create a program that requests a number and then displays the message", "The number entered was [number].")

	var i int
	fmt.Scan(&i)
	fmt.Println("The number entered was", i)

	fmt.Println("3. Create a program that requests two numbers and prints the sum.")

	var a3, b3 float32

	fmt.Scan(&a3)
	fmt.Scan(&b3)

	fmt.Println("The sum of", a3, "+", b3, "is", a3+b3)

	fmt.Println("4. Create a program that requests the 4 numbers and displays the average.")

	var a4, b4, c4, d4 float32

	fmt.Scan(&a4)
	fmt.Scan(&b4)
	fmt.Scan(&c4)
	fmt.Scan(&d4)

	fmt.Println("The average of", a4, b4, c4, d4, "is", (a4+b4+c4+d4)/4)

	fmt.Println("5. Create a program that converts meters to centimeters.")

	var meters float32

	fmt.Scan(&meters)

	fmt.Println(meters, "m corresponds to", meters*100, "centimeters.")

	fmt.Println("6. Create a program that asks how much you earn per hour and the number of hours worked in the month. Calculate and display your total salary for the month in question.")

	var salaryHours, workHours float32
	fmt.Println("Tell me how much do you gain per hour:")
	fmt.Scan(&salaryHours)
	fmt.Println("Tell me how much hours do you worked this month: ")
	fmt.Scan(&workHours)
	fmt.Println("Your month salary is $", salaryHours*workHours)

	fmt.Println("7. Create a program that requests the temperature in degrees Fahrenheit, converts it and displays it in degrees Celsius.")

	var degreesF float32
	fmt.Println("Inform the temperature in Farenheit:")
	fmt.Scan(&degreesF)

	fmt.Println(degreesF, "Farenheit corresponds to", (5 * (degreesF - 32) / 9), "Celsius")

	fmt.Println("8. Create a program that calculates a person's ideal weight using the person's height as input data.")

	var height float32
	fmt.Println("Inform your height:")
	fmt.Scan(&height)
	fmt.Println("Your ideal weight is:", (72.7*height)-58, "kg")

	fmt.Println("9. Create a program that calculates a person's ideal weight, using the person's height (h) as input.")

	fmt.Println("Inform your height:")
	fmt.Scan(&height)
	fmt.Println("Your ideal weight is:", (72.7*height)-58, "kg if you are a man", (62.1*height)-44.7, "kg if you are woman")

	fmt.Println("10. Create a program that asks how much you earn per hour and the number of hours worked in the month.")
	fmt.Println("Calculate and display your total salary for the month, knowing that 11% is deducted for income tax, 8% for social security and 5% for the union. Create a program that gives us:")
	fmt.Println("gross salary.")
	fmt.Println("how much you paid to social security.")
	fmt.Println("how much you paid to the union.")
	fmt.Println("net salary. Calculate the discounts and the net salary, according to the table below: + Gross Salary: R$ - IR (11%): R$ - INSS (8%): R$ - Union (5%): R$ = Net Salary: R$ Note: Gross Salary - Discounts = Net Salary. ")

	var hourSalary, hoursWorked, IR, INSS, socialsecurity, netSalary, grossSalary float32

	fmt.Println("Tell me how much do you gain per hour:")
	fmt.Scan(&hourSalary)
	fmt.Println("Tell me how much hours do you worked this month:")
	fmt.Scan(&hoursWorked)

	grossSalary = hourSalary * hoursWorked
	IR = grossSalary * (11.0 / 100)
	INSS = grossSalary * (8.0 / 100)
	socialsecurity = grossSalary * (5.0 / 100)
	netSalary = grossSalary - IR - INSS - socialsecurity

	fmt.Println("+ Gross Salary:", grossSalary)
	fmt.Println("- IR (11%):", IR)
	fmt.Println("- INSS (8%):", INSS)
	fmt.Println("- Social Security (5%):", socialsecurity)
	fmt.Println("= Net Salary: ", netSalary)

	fmt.Println("11. Create a program that asks for the size of a file to download (in MB) and the speed of an Internet link (in Mbps), calculate and inform the approximate download time of the file using this link (in minutes).")

	var archiveSize, internetSpeed, archiveSizeMegaBits float32

	fmt.Println("Enter the file size in MB (Mega Bytes):")
	fmt.Scan(&archiveSize)
	fmt.Println("Enter your internet speed in Mbps (Mega bits per second):")
	fmt.Scan(&internetSpeed)
	archiveSizeMegaBits = archiveSize * 8
	fmt.Println("The approximate time to download the file is", archiveSizeMegaBits/internetSpeed, "seconds")

	fmt.Println("12. Create a program that asks for two numbers and prints the largest of them.")

	var number1, number2 int

	fmt.Println("Enter an integer number:")
	fmt.Scan(&number1)
	fmt.Println("Enter another integer number:")
	fmt.Scan(&number2)

	if number1 > number2 {
		fmt.Println(number1)
	} else {
		fmt.Println(number2)
	}

	fmt.Println("13. Create a program that asks for a value and shows on the screen whether the value is positive or negative.")

	var value float32

	fmt.Println("Enter a value:")
	fmt.Scan(&value)

	if value > 0 {
		fmt.Println("The value entered is positive.")
	} else if value < 0 {
		fmt.Println("The value entered is negative.")
	}

	fmt.Println("14. Create a program that checks whether a typed letter is a vowel or a consonant.")

	var letter string
	fmt.Println("Enter a letter: ")
	fmt.Scan(&letter)

	letter = strings.ToUpper(letter)

	if strings.Compare(letter, "A") == 0 || strings.Compare(letter, "E") == 0 || strings.Compare(letter, "I") == 0 || strings.Compare(letter, "O") == 0 || strings.Compare(letter, "U") == 0 {
		fmt.Println("Vowel")
	} else {
		fmt.Println("Consonant")
	}

	fmt.Println("15. Create a program to read two partial grades of a student.")
	fmt.Println("The program should calculate the average achieved by each student and display:")
	fmt.Println("The message \"Approved\", if the average achieved is greater than or equal to seven;")
	fmt.Println("The message \"Failed\", if the average is less than seven;")
	fmt.Println("The message \"Approved with Distinction\", if the average is equal to ten.")

	var firstGrade, secondGrade, averageGrade float32

	fmt.Println("Enter the first grade: ")
	fmt.Scan(&firstGrade)

	fmt.Println("Enter the second grade: ")
	fmt.Scan(&secondGrade)

	averageGrade = (firstGrade + secondGrade) / 2.0

	if averageGrade == 10 {
		fmt.Println("Approved with Distinction")
	} else if averageGrade >= 7 {
		fmt.Println("Approved")
	} else {
		fmt.Println("Failed")
	}

	fmt.Println("16. Create a program that reads three numbers and displays the largest of them.")

	var numberInformed, biggerNumber float32

	fmt.Println("Enter the first number:")
	fmt.Scan(&numberInformed)
	fmt.Println("Enter the second number:")
	fmt.Scan(&biggerNumber)

	if numberInformed > biggerNumber {
		biggerNumber = numberInformed
	}

	fmt.Println("Enter the third number: ")
	fmt.Scan(&numberInformed)

	if numberInformed > biggerNumber {
		biggerNumber = numberInformed
	}

	fmt.Println(biggerNumber, "has the bigger number.")

}
