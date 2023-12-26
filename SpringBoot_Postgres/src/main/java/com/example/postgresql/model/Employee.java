package com.example.postgresql.model;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;


@Entity

public class Employee {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long employeeid;

    private String firstname;

    private String lastname;

    public Employee() {
    }

    public Employee(String firstName, String lastName) {
        this.firstname = firstName;
        this.lastname = lastName;
    }

    public Long getId() {
        return employeeid;
    }

    public void setId(Long id) {
        this.employeeid = id;
    }

    public String getFirstName() {
        return firstname;
    }

    public void setFirstName(String firstName) {
        this.firstname = firstName;
    }

    public String getLastName() {
        return lastname;
    }

    public void setLastName(String lastName) {
        this.lastname = lastName;
    }
}
