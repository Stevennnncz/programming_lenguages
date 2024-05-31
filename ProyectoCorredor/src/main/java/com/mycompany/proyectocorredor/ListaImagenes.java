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
public class ListaImagenes {
    private static ArrayList<Imagen> listaImagenes;

    public static ArrayList<Imagen> getInstance() {
        if (listaImagenes == null) {
            listaImagenes = new ArrayList<Imagen>();
        }
        return listaImagenes;
    }

    public static void addImage(Imagen imagen) {
        getInstance().add(imagen);
    }
}
