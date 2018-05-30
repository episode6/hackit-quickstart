package {{ .Group }}.main

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import {{ .Group }}.R
import {{ .Group }}.base.BaseFragment
import kotlinx.android.synthetic.main.fragment_main.*

class MainFragment : BaseFragment() {

  override fun onCreateView(
      inflater: LayoutInflater,
      container: ViewGroup?,
      savedInstanceState: Bundle?): View? {
    return inflater.inflate(R.layout.fragment_main, container, false)
  }


  override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
    tv_hello_world.text = "Hello Earth!"
  }
}
