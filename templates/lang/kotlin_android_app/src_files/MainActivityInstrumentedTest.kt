package {{ .Group }};

import android.content.Intent
import android.support.test.InstrumentationRegistry
import android.support.test.runner.AndroidJUnit4
import org.assertj.core.api.Assertions.assertThat
import org.junit.Test
import org.junit.runner.RunWith

/**
 * Instrumentation test, which will execute on an Android device.
 *
 * @see [Testing documentation](http://d.android.com/tools/testing)
 */
@RunWith(AndroidJUnit4::class)
class MainActivityInstrumentedTest {

  @Test
  @Throws(Exception::class)
  fun launchMainActivity() {
    val instrumentation = InstrumentationRegistry.getInstrumentation()
    val activityMonitor = instrumentation.addMonitor(MainActivity::class.java.name, null, false)


    val intent = Intent(Intent.ACTION_MAIN)
    intent.flags = Intent.FLAG_ACTIVITY_NEW_TASK
    intent.setClassName(instrumentation.targetContext, MainActivity::class.java.name)
    instrumentation.startActivitySync(intent)
    val activity = instrumentation.waitForMonitor(activityMonitor) as MainActivity

    assertThat(activity).isNotNull()
    assertThat(activity.title).isEqualTo("{{ .CamelName }}")
  }
}
