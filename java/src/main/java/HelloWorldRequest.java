import java.io.Serializable;

public class HelloWorldRequest implements Serializable {
    private String name;
    private String logo;
    public HelloWorldRequest(){

    }
    public HelloWorldRequest(String name, String logo) {
        this.name = name;
        this.logo = logo;
    }

    public String getLogo() {
        return logo;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setLogo(String logo) {
        this.logo = logo;
    }
}
