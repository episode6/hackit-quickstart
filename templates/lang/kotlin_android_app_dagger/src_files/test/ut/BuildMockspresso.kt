package {{ .Group }}

import {{ .Group }}.app.MockspressoTestApp
import com.episode6.hackit.mockspresso.Mockspresso
import com.episode6.hackit.mockspresso.quick.BuildQuickMockspresso

object BuildMockspresso {
  fun withDefaults(): Mockspresso.Builder {
    return BuildQuickMockspresso.with()
        .injector().dagger()
        .mocker().mockito()
  }

  fun forRobolectric(): Mockspresso.Builder {
    return MockspressoTestApp.get().buildUpon()
  }
}
