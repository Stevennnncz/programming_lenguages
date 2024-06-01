/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.proyectocorredor;

import java.util.ArrayList;

/**
 *
 * @author steve
 */
public class ListaTiempos {
    private static ArrayList<Tiempos> listaTiempos;

    public static ArrayList<Tiempos> getInstance() {
        if (listaTiempos == null) {
            listaTiempos = new ArrayList<Tiempos>();
        }
        return listaTiempos;
    }

    public static void addTiempos(Tiempos tiempo) {
        getInstance().add(tiempo);
    }
    
    
    public static ArrayList<Tiempos> TiemposPorNombre(String nombre) {
        ArrayList<Tiempos> listaTiemposCorredor = new ArrayList<Tiempos>();
        for (Tiempos tiempo : ListaTiempos.getInstance()) {
            if (tiempo.getCorredor().getNombre().equals(nombre)) { 
                listaTiemposCorredor.add(tiempo);
            }}
            return listaTiemposCorredor;
        
}
}

