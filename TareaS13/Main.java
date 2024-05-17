import java.util.ArrayList;
import java.util.List;

// Clase para representar una Persona
class Persona {
    private String nombre;
    private int edad;

    public Persona(String nombre, int edad) {
        this.nombre = nombre;
        this.edad = edad;
    }

    @Override
    public String toString() {
        return "Nombre: " + nombre + ", Edad: " + edad;
    }
}

// Clase abstracta para eventos
abstract class Evento {
    private String nombre;

    public Evento(String nombre) {
        this.nombre = nombre;
    }

    @Override
    public String toString() {
        return "Evento: " + nombre;
    }
}

// Clase para eventos simples
class EventoSimple extends Evento {
    public EventoSimple(String nombre) {
        super(nombre);
    }
}

// Clase para eventos específicos (se deja abstracta para que se implementen tipos específicos)
abstract class EventoEspecifico extends Evento {
    public EventoEspecifico(String nombre) {
        super(nombre);
    }
}

// Clase para eventos específicos de tipo Conferencia
class Conferencia extends EventoEspecifico {
    private String tema;

    public Conferencia(String nombre, String tema) {
        super(nombre);
        this.tema = tema;
    }

    @Override
    public String toString() {
        return super.toString() + ", Tema: " + tema;
    }
}

// Clase abstracta para contactos
abstract class Contacto {
    private Persona persona;

    public Contacto(Persona persona) {
        this.persona = persona;
    }

    @Override
    public String toString() {
        return persona.toString();
    }
}

// Clase para contactos simples
class ContactoSimple extends Contacto {
    public ContactoSimple(Persona persona) {
        super(persona);
    }
}

// Clase para contactos específicos (se deja abstracta para que se implementen tipos específicos)
abstract class ContactoEspecifico extends Contacto {
    public ContactoEspecifico(Persona persona) {
        super(persona);
    }
}

// Clase para contactos específicos de tipo ContactoFamiliar
class ContactoFamiliar extends ContactoEspecifico {
    private String relacion;

    public ContactoFamiliar(Persona persona, String relacion) {
        super(persona);
        this.relacion = relacion;
    }

    @Override
    public String toString() {
        return super.toString() + ", Relación: " + relacion;
    }
}

// Clase para contactos específicos de tipo ContactoEmpresarial
class ContactoEmpresarial extends ContactoEspecifico {
    private String empresa;

    public ContactoEmpresarial(Persona persona, String empresa) {
        super(persona);
        this.empresa = empresa;
    }

    @Override
    public String toString() {
        return super.toString() + ", Empresa: " + empresa;
    }
}

// Clase para manejar la agenda
class Agenda {
    private static Agenda instanciaUnica; // Instancia única para Singleton
    private List<Object> elementos;

    private Agenda() {
        elementos = new ArrayList<>();
    }

    // Método estático para obtener la instancia única (Lazy Singleton)
    public static Agenda obtenerInstancia() {
        if (instanciaUnica == null) {
            instanciaUnica = new Agenda();
        }
        return instanciaUnica;
    }

    // Método para añadir un elemento a la agenda
    public void agregarElemento(Object elemento) {
        elementos.add(elemento);
    }

    // Método para eliminar un elemento de la agenda
    public void eliminarElemento(Object elemento) {
        elementos.remove(elemento);
    }

    // Método para modificar un elemento de la agenda
    public void modificarElemento(Object elementoAntiguo, Object elementoNuevo) {
        int indice = elementos.indexOf(elementoAntiguo);
        if (indice != -1) {
            elementos.set(indice, elementoNuevo);
        }
    }

    // Método para filtrar y mostrar eventos de la agenda
    public void mostrarEventos() {
        elementos.stream()
                 .filter(obj -> obj instanceof Evento)
                 .forEach(System.out::println);
    }

    // Método para filtrar y mostrar contactos de la agenda
    public void mostrarContactos() {
        elementos.stream()
                 .filter(obj -> obj instanceof Contacto)
                 .forEach(System.out::println);
    }
}

public class Main {
    public static void main(String[] args) {
        // Crear instancias de personas
        Persona persona1 = new Persona("Juan", 30);
        Persona persona2 = new Persona("Maria", 25);

        // Crear instancias de contactos
        ContactoSimple contactoSimple = new ContactoSimple(persona1);
        ContactoFamiliar contactoFamiliar = new ContactoFamiliar(persona2, "Hermano");
        ContactoEmpresarial contactoEmpresarial = new ContactoEmpresarial(persona1, "Empresa A");

        // Crear instancias de eventos
        EventoSimple eventoSimple = new EventoSimple("Fiesta");
        Conferencia conferencia = new Conferencia("Conferencia de Tecnología", "Inteligencia Artificial");

        // Obtener instancia de la agenda (Singleton)
        Agenda agenda = Agenda.obtenerInstancia();

        // Agregar elementos a la agenda
        agenda.agregarElemento(contactoSimple);
        agenda.agregarElemento(contactoFamiliar);
        agenda.agregarElemento(contactoEmpresarial);
        agenda.agregarElemento(eventoSimple);
        agenda.agregarElemento(conferencia);

        // Mostrar eventos
        System.out.println("Eventos:");
        agenda.mostrarEventos();

        // Mostrar contactos
        System.out.println("\nContactos:");
        agenda.mostrarContactos();
    }
}


//Justificación de los patrones de diseño utilizados:

//Singleton: Se utiliza el patrón Singleton para garantizar que solo haya una instancia de la clase Agenda en toda la aplicación. Esto es útil en situaciones donde se necesita una única fuente de verdad para la gestión de datos, como una agenda en este caso. Se implementó un Singleton "Lazy" para garantizar la creación de la instancia solo cuando sea necesario, lo que mejora el rendimiento si la instancia no se utiliza en cada ejecución.

//Abstract Factory: No se utilizó el patrón Abstract Factory en esta implementación porque no hay una clara necesidad de crear familias de objetos relacionados. El problema se centra más en la gestión de elementos individuales (contactos y eventos) en la agenda. El uso de Abstract Factory sería apropiado si hubiera diferentes tipos de agendas o fabricantes de contactos/eventos. En este caso, no hay tales requerimientos, por lo que su uso sería innecesario y agregaría complejidad innecesaria al diseño.