import java.time.LocalDate;

public class Student {
    private String name;
    private LocalDate birth;

    private Student() {} // preventing direct instantiation

    public static class Builder {
        private Student student = new Student();

        public Builder name(String name) {
            student.name = name;
            return this;
        }
        public Builder birth(LocalDate date) {
            student.birth = date;
            return this;
        }
        public Student build() {
            if (student.name == null
                    || "".equals(student.name.trim())) {
                throw new RuntimeException(
                        "Student needs a name");
            }
            return student;
        }
    }

    public static void main() {
        Student student = new Student.Builder()
                .name("John")
                .birth(LocalDate.of(1979,5,25))
                .build();
    }
}