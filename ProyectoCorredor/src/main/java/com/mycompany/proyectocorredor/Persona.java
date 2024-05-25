/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.proyectocorredor;

/**
 *
 * @author steve
 */
public class Persona {
    private String name;
    private String mail;
    private String phone;
    private Boolean type;
    
    public Persona(String n, String m, String p, Boolean b) {
        this.name = n;
        this.mail = m;
        this.phone= p;
        this.type = b;
    }

    @Override
    public String toString() {
        return "Nombre: " + name + ", Correo: " + mail + ", Telefono: " + phone + ", Tipo de sangre: " + type;
    }
}
