package {{ .Group }};

import android.app.Instrumentation;
import android.content.Intent;
import android.support.test.InstrumentationRegistry;
import android.support.test.runner.AndroidJUnit4;

import {{ .Group }}.main.MainActivity;

import org.junit.Test;
import org.junit.runner.RunWith;

import static org.assertj.core.api.Assertions.assertThat;


/**
 * Instrumentation test, which will execute on an Android device.
 *
 * @see <a href="http://d.android.com/tools/testing">Testing documentation</a>
 */
@RunWith(AndroidJUnit4.class)
public class MainActivityInstrumentedTest {

  @Test
  public void launchMainActivity() throws Exception {
    Instrumentation instrumentation = InstrumentationRegistry.getInstrumentation();
    Instrumentation.ActivityMonitor activityMonitor =
        instrumentation.addMonitor(MainActivity.class.getName(), null, false);


    Intent intent = new Intent(Intent.ACTION_MAIN);
    intent.setFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
    intent.setClassName(instrumentation.getTargetContext(), MainActivity.class.getName());
    instrumentation.startActivitySync(intent);
    MainActivity activity = (MainActivity) instrumentation.waitForMonitor(activityMonitor);

    assertThat(activity).isNotNull();
    assertThat(activity.getTitle()).isEqualTo("{{ .CamelName }}");
  }
}
