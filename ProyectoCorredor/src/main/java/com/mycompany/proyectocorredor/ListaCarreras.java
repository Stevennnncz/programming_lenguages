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
public class ListaCarreras {
    private static ArrayList<Carrera> listaCarreras;

    public static ArrayList<Carrera> getInstance() {
        if (listaCarreras == null) {
            listaCarreras = new ArrayList<Carrera>();
        }
        return listaCarreras;
    }

    public static void addCarrera(Carrera carrera) {
        getInstance().add(carrera);
    }
    
}





    
