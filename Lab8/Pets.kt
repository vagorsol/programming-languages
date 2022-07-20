/**
* Setting up the Pet UMl structure.
* @author ayang
* Created: October 25, 2021
 */

open class Pet(val id: String, val name: String) {
    open fun sound(): String {
        return "*insert generic animal sound here*"
    }

    override fun toString() = String.format("%s: %s", id, name)
}

class Cat: Pet {
    val breed: String
    val hairLength: Double

    constructor(id: String, name: String, inBreed: String, inHairLength: Double): super(id, name){
        breed = inBreed
        hairLength = inHairLength
    }

    override fun sound(): String {
        return "Meow!"
    }

    override fun toString() = String.format("%s: %s, %s, %f", id, name, breed, hairLength)
}

open class Dog: Pet {
    val group: String
    val breed: String
    val hairLength: Double
    val doubleCoat: Boolean

    constructor(id: String, name: String, inGroup: String, inBreed: String, inHairLength: Double, inDoubleCoat: Boolean): super(id, name){
        group = inGroup
        breed = inBreed
        hairLength = inHairLength
        doubleCoat = inDoubleCoat
    }

    override fun toString() = String.format("%s: %s, %s, %s, %f, %b", id, name, group, breed, hairLength, doubleCoat)

    override fun sound(): String {
        return "Woof!"
    }
}

class WorkingDog: Dog {
    val typeOfWork: String

    // dog is the super class
    constructor(id: String, name: String, group: String, breed: String, hairLength: Double, doubleCoat: Boolean, inTypeOfWork: String): super(id, name, group, breed, hairLength, doubleCoat){
        typeOfWork = inTypeOfWork
    }

    override fun toString() = String.format("%s: %s, %s, %s, %f, %b, %s", id, name, group, breed, hairLength, doubleCoat, typeOfWork)

    override fun sound(): String{
        return "Dignified Howl"
    }
}

fun main(args:Array<String>) {
    println("Pet")
    val pet = Pet("1978", "Milo")
    println(pet)
    println(pet.sound())
    
    println("Cat")
    val cat = Cat("8087", "Louie", "Short Haired", 4.20)
    println(cat)
    println(cat.sound())
    
    println("Dog")
    val dog = Dog("0843", "Toby", "Toy", "Pomeranian", 10.10, true)
    println(dog)
    println(dog.sound())
    
    println("Working Dog")
    val workingDog = WorkingDog("8876", "Kimchi", "Herding", "Corgi", 0.69, true, "Surgery Assistant")
    println(workingDog)
    println(workingDog.sound())

    // show dynamic dispatch
    // cast up, but if you call it calls the working dog one (the one it is most recently defined as)
    // static dispatch: would have printed out the toString for pet (since it is what it is defined as)
    val tucker: Pet = workingDog
    println(tucker)
}