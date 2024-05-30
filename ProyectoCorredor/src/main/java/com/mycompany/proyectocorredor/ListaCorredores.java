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
public class ListaCorredores {
    private static ArrayList<Corredor> listaCorredor;

    public static ArrayList<Corredor> getInstance() {
        if (listaCorredor == null) {
            listaCorredor = new ArrayList<Corredor>();
        }
        return listaCorredor;
    }

    public static void addRunner(Corredor corredor) {
        getInstance().add(corredor);
    }

    public static Corredor CorredorPorNombre(String nombre) {
        for (Corredor corredor : getInstance()) {
            if (corredor.getNombre().equals(nombre)) {
                return corredor;
            }
        }
        return null;
    }
}




    
