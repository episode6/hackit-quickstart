package {{ .Group }}

import {{ .Group }}.app.MockspressoTestApp
import com.episode6.hackit.mockspresso.quick.QuickMockspresso
import com.episode6.hackit.mockspresso.quick.BuildQuickMockspresso

object BuildMockspresso {
  fun withDefaults(): QuickMockspresso.Builder {
    return BuildQuickMockspresso.with()
        .injector().dagger()
        .mocker().mockito()
  }

  fun forRobolectric(): QuickMockspresso.Builder {
    return MockspressoTestApp.get().buildUpon()
  }
}
