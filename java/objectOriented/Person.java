package objectOriented;

public class Person {
    protected String name;// private的字段子类无法访问
    protected int age;

    public Person() {

    }

    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }

    public void setName() {
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getName() {
        return "person:" + this.name;
    }
}