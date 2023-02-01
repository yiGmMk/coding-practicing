package objectOriented;

public class Person {
    protected String name;// private的字段子类无法访问
    protected int age;
    private int cash;

    public Person() {

    }

    public Person(String name, int age, int cash) {
        this.name = name;
        this.age = age;
        this.cash = cash;
    }

    public void setName() {
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setCash(int cash) {
        this.cash = cash;
    }

    public String getName() {
        return "person:" + this.name;
    }
}