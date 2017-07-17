package {{ .Group }};

import com.episode6.hackit.mockspresso.Mockspresso;

import {{ .Group }}.app.MockspressoTestApp;

import org.junit.Before;

import static com.episode6.hackit.mockspresso.BuildMockspresso.with;

public class BuildMockspresso {
  public static Mockspresso.Builder withDefaults() {
    return with()
        .injector().dagger()
        .mocker().mockito();
  }

  public static Mockspresso.Builder forRobolectric() {
    return MockspressoTestApp.get().applicationMockspresso().buildUpon()
        .testResources(new Object() {
          @Before
          public void setup(Mockspresso mockspresso) {
            MockspressoTestApp.get().injectComponentsUsing(mockspresso);
          }
        });
  }
}
