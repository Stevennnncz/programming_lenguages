/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.proyectocorredor;
import java.util.ArrayList;
import com.mycompany.proyectocorredor.ListaCorredores;
/**
 *
 * @author steve
 */
public class ProyectoCorredor {
    
    static ArrayList<Persona> listaAdmins = new ArrayList<>();
    
    public static void main(String[] args){
    Login ventana = new Login();
    ventana.setVisible(true);
    
    
    Corredor corredor1 = new Corredor ("Leiner", "Leiner@gmail.com", "123", "a+");
    Corredor corredor2 = new Corredor ("Steven", "Steven@gmail.com", "234", "b+");
    
    ListaCorredores.addRunner(corredor1);
    ListaCorredores.addRunner(corredor2);
    
    Persona persona3 = new Persona("Karina", "Karina@gmail.com", "1");
    Persona persona4 = new Persona("Walter", "Walter@gmail.com", "1");
    
    listaAdmins.add(persona3);
    listaAdmins.add(persona4);
    
    
    }
    
    public static void login(String correo, String tipo) {
    if ("Corredor".equals(tipo)) {
        for (Corredor corredor : ListaCorredores.getInstance()) {
            if (correo.equals(corredor.getCorreo())) {
                menuCorredor ventanaC = new menuCorredor();
                ventanaC.setVisible(true);
                return;
            }
        }
    } else if ("Administrador".equals(tipo)) {
        for (Persona persona : listaAdmins) {
            if (correo.equals(persona.getCorreo())) {
                menuAdmin ventanaA = new menuAdmin();
                ventanaA.setVisible(true);
                return;
            }
        }
    }
}}
        
    
    
