package objectOriented;

public class Student extends Person {
    public double score;

    @Override
    public void setName(String name) {
        this.name = "student:" + name;
    }

    public Student() {

    }

    public Student(String name, int age, int score) {
        super(name, age);
        this.score = score;
    }

    @Override
    public String getName() {
        return "student:" + this.name;
    }
}