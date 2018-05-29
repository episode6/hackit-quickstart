package {{ .Group }};

import com.episode6.hackit.mockspresso.Mockspresso;

import {{ .Group }}.app.MockspressoTestApp;

import static com.episode6.hackit.mockspresso.quick.QuickBuildMockspresso.with;

public class BuildMockspresso {
  public static Mockspresso.Builder withDefaults() {
    return with()
        .injector().dagger()
        .mocker().mockito();
  }

  public static Mockspresso.Builder forRobolectric() {
    return MockspressoTestApp.get().buildUpon();
  }
}
