/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.proyectocorredor;

/**
 *
 * @author steve
 */
public class Tiempos {
    private String corredor;
    private  int tiempoSeg;
    private int tiempoTotal;
    private String posicionFinal;
    
    
    public Tiempos(String c, int ts, int tt, String pf){
        this.corredor = c;
        this.tiempoSeg= ts;
        this.tiempoTotal=tt;
        this.posicionFinal = pf;
    }

}
