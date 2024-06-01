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
    private Corredor corredor;
    private  String tiempoSeg;
    private String tiempoTotal;
    
    
    public Tiempos(Corredor c, String ts, String tt){
        this.corredor = c;
        this.tiempoSeg= ts;
        this.tiempoTotal=tt;
    }

    public Corredor getCorredor() {
        return corredor;
    }

    public String getTiempoSeg() {
        return tiempoSeg;
    }

    public String getTiempoTotal() {
        return tiempoTotal;
    }

        @Override
    public String toString() {
        return "Tiempos{" + "corredor=" + corredor + ", tiempoSeg=" + tiempoSeg + ", tiempoTotal=" + tiempoTotal + '}';
    }
}
