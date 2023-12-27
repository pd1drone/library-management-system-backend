CREATE DATABASE lms;

USE lms;

CREATE TABLE `Admin` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `Username` varchar(255),
  `Password` varchar(255)
);

CREATE TABLE `Students` (
  `StudentID` int PRIMARY KEY AUTO_INCREMENT,
  `FirstName` varchar(255),
  `MiddleName` varchar(255),
  `LastName` varchar(255),
  `DateOfBirth` varchar(255),
  `Gender` varchar(255),
  `Address` varchar(255),
  `Email` varchar(255),
  `PhoneNumber` varchar(255),
  `RegistrationDate` varchar(255),
  `LibraryCardNumber` varchar(255)
);

CREATE TABLE `Books` (
  `BookID` int PRIMARY KEY AUTO_INCREMENT,
  `BookShelveAddress` varchar(255),
  `Title` varchar(255),
  `Author` varchar(255),
  `ISBN` varchar(255),
  `Genre` varchar(255),
  `PublicationDate` varchar(255),
  `Publisher` varchar(255),
  `Description` varchar(255),
  `IsAvailable` bool
);

CREATE TABLE `BorrowedBooks` (
  `TransactionID` int PRIMARY KEY AUTO_INCREMENT,
  `StudentID` int,
  `BookID` int,
  `BorrowedDate` int,
  `ReturnedDate` int,
  `OverdueDate` int,
  `IsReturned` bool
);


-- username= admin , password = admin123
INSERT INTO Admin (Username, Password) VALUES ('admin','0192023a7bbd73250516f069df18b500');