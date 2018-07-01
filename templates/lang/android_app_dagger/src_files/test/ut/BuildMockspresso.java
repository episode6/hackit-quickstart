package {{ .Group }};

import com.episode6.hackit.mockspresso.quick.QuickMockspresso;

import {{ .Group }}.app.MockspressoTestApp;

import static com.episode6.hackit.mockspresso.quick.QuickBuildMockspresso.with;

public class BuildMockspresso {
  public static QuickMockspresso.Builder withDefaults() {
    return with()
        .injector().dagger()
        .mocker().mockito();
  }

  public static QuickMockspresso.Builder forRobolectric() {
    return MockspressoTestApp.get().buildUpon();
  }
}
