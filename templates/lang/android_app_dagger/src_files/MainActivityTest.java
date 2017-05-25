package {{ .Group }};

import org.junit.Test;
import org.junit.runner.RunWith;
import org.robolectric.Robolectric;
import org.robolectric.RobolectricTestRunner;
import org.robolectric.annotation.Config;

import {{ .Group }}.main.MainActivity;

import static org.fest.assertions.api.Assertions.assertThat;

@RunWith(RobolectricTestRunner.class)
@Config(constants = BuildConfig.class)
public class MainActivityTest {

  @Test
  public void placeholderTest() {
    MainActivity activity = Robolectric.setupActivity(MainActivity.class);

    assertThat(activity.getTitle())
        .isEqualTo("{{ .CamelName }}");
  }
}
