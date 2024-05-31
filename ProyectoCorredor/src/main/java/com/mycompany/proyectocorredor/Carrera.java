/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.proyectocorredor;

/**
 *
 * @author steve
 */
public class Carrera {
    private Corredor corredor;
    private String distance;
    private String laterality;
    private String contactoEmergencia;
    private String category;
    private int nCorredor;
    
    public Carrera(Corredor co, String d, String l, String ce, String c) {
        this.corredor = co;
        this.distance = d;
        this.laterality = l;
        this.contactoEmergencia = ce;
        this.category = c;
    }

    public Corredor getCorredor(){
    return corredor;
}
 
    public void setIdCorredor(int x){
    this.nCorredor = x;
}
        
}
