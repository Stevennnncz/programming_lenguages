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
    private String nombre;
    private String correo;
    private String telefono;
    
    
    public Persona(String n, String c, String t) {
        this.nombre = n;
        this.correo = c;
        this.telefono = t;
        
    }
    
    public String getNombre(){
    return nombre;
    }
    
    public String getCorreo(){
    return correo;
    }
    
    @Override
    public String toString() {
        return "Nombre: " + nombre + ", Correo: " + correo + ", Telefono: " + telefono;
    }
}
