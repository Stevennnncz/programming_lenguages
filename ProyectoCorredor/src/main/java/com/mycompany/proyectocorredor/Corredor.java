/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.proyectocorredor;

/**
 *
 * @author steve
 */
public class Corredor extends Persona {
      
        private String tipoSangre;
    public Corredor(String nombre, String correo, String telefono, String  ts) { 
        super(nombre,correo,telefono);
        this.tipoSangre = ts;
    }
    
    public String getTipoSangre(){
        return tipoSangre;
    }

}
