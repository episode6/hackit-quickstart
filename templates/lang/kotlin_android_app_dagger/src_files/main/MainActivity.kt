package {{ .Group }}.main

import android.os.Bundle
import {{ .Group }}.R
import {{ .Group }}.base.BaseAppCompatActivity

class MainActivity : BaseAppCompatActivity() {

  override fun onCreate(savedInstanceState: Bundle?) {
    super.onCreate(savedInstanceState)
    setContentView(R.layout.activity_main)

    if (supportFragmentManager.findFragmentById(R.id.main_content) == null) {
      supportFragmentManager.beginTransaction()
          .add(R.id.main_content, MainFragment())
          .commit()
    }
  }
}
